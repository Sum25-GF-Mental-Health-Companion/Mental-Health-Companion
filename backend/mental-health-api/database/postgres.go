package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB

// InitDatabase connects to PostgreSQL using pgx + database/sql
func InitDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname)

	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	if err = DB.PingContext(context.Background()); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	log.Println("Connected to PostgreSQL using pgx/sqlc")

	RunMigrations(DB, "./migrations")
}

func RunMigrations(db *sql.DB, migrationsDir string) {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Goose dialect error: %v", err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
