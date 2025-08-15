package handlers

import "net/http"

// GetMyInventory implements apigen.ServerInterface.
func (s Server) GetMyInventory(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetPlayerInventory implements apigen.ServerInterface.
func (s Server) GetPlayerInventory(w http.ResponseWriter, r *http.Request, playerId string) {
	panic("unimplemented")
}
