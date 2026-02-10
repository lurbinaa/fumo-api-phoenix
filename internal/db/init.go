package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool       *pgxpool.Pool
	FumosRepo  *FumosRepo
	UsersRepo  *UsersRepo
	TokensRepo *TokensRepo
}

func NewDb() *Database {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Could not connect to dabatase: %s\n", err)
	}

	err = initSchemas(ctx, pool)
	if err != nil {

	}

	fumosRepo := newFumosRepo(pool)
	usersRepo := newUsersRepo(pool)
	tokensRepo := newTokensRepo(pool)

	return &Database{
		Pool:       pool,
		FumosRepo:  fumosRepo,
		UsersRepo:  usersRepo,
		TokensRepo: tokensRepo,
	}
}

func initSchemas(ctx context.Context, pool *pgxpool.Pool) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("Transaction init fail: %w", err)
	}
	defer tx.Rollback(ctx)

	var queries = []string{
		`DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
				CREATE TYPE user_role AS ENUM ('superadmin','admin','user');
			END IF;
		END $$;`,
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			discord_id TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			role user_role NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS tokens (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			token_hash TEXT UNIQUE NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			revoked_at TIMESTAMPTZ NULL
		);`,
		`CREATE TABLE IF NOT EXISTS fumo_media (
			id SERIAL PRIMARY KEY,
			fumo_ids INT[] NOT NULL,
			description TEXT NULL,
			media_url TEXT NOT NULL,
			approved BOOLEAN NOT NULL DEFAULT FALSE
		);`,
	}
	for i, q := range queries {
		_, err := tx.Exec(ctx, q)
		if err != nil {
			return fmt.Errorf("Error at table #%d: %w", i, err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("Commit fail: %w", err)
	}

	return nil
}
