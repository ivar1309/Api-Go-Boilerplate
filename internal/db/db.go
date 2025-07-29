package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}

	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
    		id SERIAL PRIMARY KEY,
    		username VARCHAR(50) UNIQUE NOT NULL,
    		password_hash TEXT NOT NULL,
    		role VARCHAR(20) NOT NULL DEFAULT 'user'
		)
	`
	_, err = pool.Exec(context.Background(), createUserTable)

	DB = pool
}

func GetDB() *pgxpool.Pool {
	return DB
}

func CloseDB() {
	DB.Close()
}
