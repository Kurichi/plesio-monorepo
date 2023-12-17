package database

import (
	"database/sql"
	"fmt"

	"github.com/Kurichi/plesio-monorepo/services/item/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	*bun.DB
}

func New(cfg *config.DBConfig) *DB {
	cfg.Host = "localhost"
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return &DB{db}
}
