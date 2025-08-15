package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	apigen "github.com/yazmeyaa/duck_in_the_sea/internal/api/http_server/generated"
	usecase "github.com/yazmeyaa/duck_in_the_sea/internal/users/use_case"
)

type Server struct {
	logger *slog.Logger
	uc     usecase.UserUseCase
}

// GetPlayer implements apigen.ServerInterface.
func (s Server) GetPlayer(w http.ResponseWriter, r *http.Request, params apigen.GetPlayerParams) {
	panic("unimplemented")
}

func (s *Server) writeErrorJSON(w http.ResponseWriter, code int, errMsg string) {
	if err := json.NewEncoder(w).Encode(apigen.DefaultErrorResponse{Message: errMsg}); err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func (s *Server) writeErrorJSONInternal(w http.ResponseWriter, errMsg string) {
	s.writeErrorJSON(w, http.StatusInternalServerError, errMsg)
}

// AuthViaTelegram implements apigen.ServerInterface.
func (s Server) AuthViaTelegram(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// RefreshJwtToken implements apigen.ServerInterface.
func (s Server) RefreshJwtToken(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

var _ apigen.ServerInterface = Server{}

func NewServer(uc usecase.UserUseCase, logger *slog.Logger) *Server {
	return &Server{
		uc:     uc,
		logger: logger,
	}
}
