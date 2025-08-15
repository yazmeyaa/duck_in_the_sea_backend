package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	apigen "github.com/yazmeyaa/duck_in_the_sea/internal/api/http_server/generated"
	"github.com/yazmeyaa/duck_in_the_sea/internal/users/port"
)

// CreateUser implements apigen.ServerInterface.
func (s Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body := apigen.CreateUserJSONBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		s.logger.Error("Error while reading request body", slog.Any("error", err))
		s.writeErrorJSON(w, http.StatusBadRequest, "Cannot decode JSON body")
		return
	}
	data := port.CreateUserDTO{
		Name:       body.Name,
		TelegramID: body.TelegramId,
	}

	id, err := s.uc.CreateUser(r.Context(), data)
	if err != nil {
		s.logger.Error("Error while creating user", slog.Any("error", err))
		s.writeErrorJSON(w, http.StatusBadRequest, "Error whiel creating user")
		return
	}

	user, err := s.uc.GetUserByID(r.Context(), id)
	if err != nil {
		s.logger.Error("Error while retrieving user", slog.Any("error", err))
		s.writeErrorJSONInternal(w, "Internal error")
		return
	}

	res := apigen.CreateUserResponse{
		CreatedAt:  user.CreatedAt,
		Id:         user.ID,
		Name:       user.Name,
		TelegramId: user.TelegramID,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		s.logger.Error("Error while serializing user structure", slog.Any("error", err))
		s.writeErrorJSONInternal(w, "Internal server error. Encoder error.")
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetUserById implements apigen.ServerInterface.
func (s Server) GetUserById(w http.ResponseWriter, r *http.Request, userId int) {
	user, err := s.uc.GetUserByID(r.Context(), userId)
	if err != nil {
		s.writeErrorJSON(w, http.StatusBadRequest, "Error while processing request")
		return
	}

	response := apigen.GetUserByIdResponse{
		CreatedAt:  user.CreatedAt,
		Id:         user.ID,
		Name:       user.Name,
		TelegramId: user.TelegramID,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		s.logger.Error("Error while writing JSON user model", slog.Any("error", err))
		s.writeErrorJSONInternal(w, "Unable to write response")
		return
	}
}
