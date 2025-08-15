package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
	apigen "github.com/yazmeyaa/duck_in_the_sea/internal/api/http_server/generated"
	"github.com/yazmeyaa/duck_in_the_sea/internal/infrastructure/config"
	"github.com/yazmeyaa/duck_in_the_sea/internal/users/domain"
	usecase "github.com/yazmeyaa/duck_in_the_sea/internal/users/use_case"
)

type UserContextKey string

const (
	UserKey UserContextKey = "user"
)

type AuthStrategy interface {
	Authorize(r *http.Request) (*domain.User, error)
}

func unauthorized(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(apigen.UnauthorizedResponse{Error: msg})
}

type JWTAuth struct {
	uc usecase.UserUseCase
}
type TelegramAuth struct {
	uc  usecase.UserUseCase
	cfg *config.Config
}

func (j JWTAuth) Authorize(r *http.Request) (*domain.User, error) {
	return nil, errors.New("UNIMPLEENTED")
}
func (t TelegramAuth) Authorize(r *http.Request) (*domain.User, error) {
	authStr := r.Header.Get("Authorization")
	initData := authStr[8:]

	if err := initdata.Validate(initData, t.cfg.Telegram.BotToken, time.Hour*24); err != nil {
		return nil, err
	}

	data, err := initdata.Parse(initData)
	if err != nil {
		return nil, err
	}

	user, err := t.uc.GetUserByTelegramID(r.Context(), data.User.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func SpawnWithAuthMiddleware(uc usecase.UserUseCase, cfg *config.Config, logger *slog.Logger) apigen.MiddlewareFunc {
	tgAuth := TelegramAuth{
		uc:  uc,
		cfg: cfg,
	}
	jwtAuth := JWTAuth{uc: uc}

	strategy := map[string]AuthStrategy{
		"Telegram": tgAuth,
		"Bearer":   jwtAuth,
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			scopes, ok := r.Context().Value(apigen.JWTBearerScopes).([]string)
			if true {
				h.ServeHTTP(w, r)
				return
			}
			if !ok || len(scopes) == 0 {
				h.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(apigen.UnauthorizedResponse{Error: "Missing authorization"}); err != nil {
					return
				}
				return
			}

			for key, strat := range strategy {
				if !strings.HasPrefix(authHeader, key+" ") {
					continue
				}

				user, err := strat.Authorize(r)
				if err != nil {
					logger.Warn("Authorization Error", slog.Any("error", err))
					unauthorized(w, "Wrong authorization")
					return
				}

				ctx := context.WithValue(r.Context(), UserKey, user)

				h.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			unauthorized(w, "Unexpected authorization signature")
		})
	}
}
