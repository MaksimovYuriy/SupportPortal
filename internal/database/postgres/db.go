package postgres

import (
	"database/sql"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB(cfg config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
