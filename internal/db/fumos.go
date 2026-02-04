package db

import "github.com/jackc/pgx/v5/pgxpool"

type FumosRepo struct {
	pool *pgxpool.Pool
}

type FumoData struct {
	Name        string
	Description *string
	MediaURL    string
	Approved    bool
}

func newFumosRepo(p *pgxpool.Pool) *FumosRepo {
	return &FumosRepo{
		pool: p,
	}
}
