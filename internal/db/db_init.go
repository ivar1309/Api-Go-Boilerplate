package db

import (
	"context"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool
var Q *Queries

func InitDB() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}

	m, err := migrate.New("file://sql/migrations", os.Getenv("MIGRATION_URL"))
	if err != nil {
		log.Printf("Database Migration Connection: %v", err)
	}
	if err := m.Up(); err != nil {
		log.Printf("Database Migration: %v", err)
	}

	Q = New(pool)
	DB = pool
}

func GetDB() *pgxpool.Pool {
	return DB
}

func GetQ() *Queries {
	return Q
}

func CloseDB() {
	DB.Close()
}
