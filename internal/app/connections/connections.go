package connections

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"medtracker/medtracker/internal/app/config"
)

type Connections struct {
	DB *sql.DB
}

func NewConnections(cfg *config.Config) (*Connections, error) {
	db, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to MySQL: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping MySQL: %w", err)
	}
	return &Connections{DB: db}, nil
}
func (c *Connections) Close() {
	if c.DB != nil {
		_ = c.DB.Close()
	}
}
