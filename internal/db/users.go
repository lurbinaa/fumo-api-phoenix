package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepo struct {
	pool *pgxpool.Pool
}

type UserRole string

const (
	UserRoleSuperadmin UserRole = "superadmin"
	UserRoleAdmin      UserRole = "admin"
	UserRoleUser       UserRole = "user"
)

type UserDataRegister struct {
	DiscordID string
	Name      string
	Role      *UserRole
}

type UserData struct {
	DiscordID string
	Name      string
	Role      UserRole
}

func newUsersRepo(p *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{
		pool: p,
	}
}

func (ur *UsersRepo) registerUser(ctx context.Context, userData UserDataRegister) error {
	role := UserRoleUser
	if userData.Role != nil {
		role = *userData.Role
	}

	_, err := ur.pool.Exec(
		ctx,
		`INSERT INTO  users (discord_id, name, role)
		VALUES ($1, $2, $3)`,
		userData.DiscordID,
		userData.Name,
		role,
	)
	if err != nil {
		return fmt.Errorf(
			"Failed to register user \"%s\": %w",
			userData.DiscordID,
			err,
		)
	}

	return nil
}
