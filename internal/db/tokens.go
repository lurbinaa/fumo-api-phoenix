package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TokensRepo struct {
	pool *pgxpool.Pool
}

func newTokensRepo(p *pgxpool.Pool) *TokensRepo {
	return &TokensRepo{pool: p}
}

func (tr *TokensRepo) StoreToken(ctx context.Context, userID int, tokenHash string) error {
	_, err := tr.pool.Exec(
		ctx,
		`INSERT INTO tokens (user_id, token_hash)
		VALUES ($1, $2)`,
		userID,
		tokenHash,
	)
	return err
}

func (tr *TokensRepo) ValidateToken(ctx context.Context, tokenHash string) (bool, error) {
	var exists bool
	err := tr.pool.QueryRow(
		ctx,
		`SELECT EXISTS(
            		SELECT 1
            		FROM tokens
            		WHERE token_hash=$1
              			AND revoked_at IS NULL)`,
		tokenHash,
	).Scan(&exists)
	return exists, err
}

func (tr *TokensRepo) RevokeToken(ctx context.Context, tokenHash string) error {
	_, err := tr.pool.Exec(
		ctx,
		`UPDATE tokens
        	SET revoked_at = NOW()
		WHERE token_hash = $1`,
		tokenHash,
	)
	return err
}
