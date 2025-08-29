package store

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/navid-fn/shorty/internal/utils"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error) {
	config, err := utils.LoadConfig()

	dbConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBusername, config.DBpassword, config.DBname, config.DBport)
	if err != nil {
		return nil, fmt.Errorf("db Open: %w", err)
	}

	db, err := sql.Open("pgx", dbConfig)
	if err != nil {
		return nil, fmt.Errorf("db Open: %w", err)
	}
	fmt.Println("Database connected...")
	return db, nil
}

func MigrateFS(db *sql.DB, migrationsFs fs.FS, dir string) error {
	goose.SetBaseFS(migrationsFs)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrations(db, dir)
}

func Migrations(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("Migrate: %w", err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose up: %w", err)
	}
	return nil
}
