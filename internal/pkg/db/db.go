package db

import (
	"database/sql"
	"fmt"
	"go-micro/internal/pkg/env"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB

func DB() *bun.DB {
	if db != nil {
		return db
	}

	username := env.ReadString("POSTGRES_USER", "admin")
	password := env.ReadString("POSTGRES_PASSWORD", "admin")
	dbname := env.ReadString("POSTGRES_DB", "go-micro")
	dsn := fmt.Sprintf("postgres://%s:%s@postgres:5432/%s?sslmode=disable", username, password, dbname)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())
	return db
}
