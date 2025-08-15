package repository

import "github.com/jackc/pgx/v4"

type resourcesManager struct {
	conn *pgx.Conn
	// Add cache
}

func NewManager(conn *pgx.Conn) *resourcesManager {
	return &resourcesManager{
		conn: conn,
	}
}
