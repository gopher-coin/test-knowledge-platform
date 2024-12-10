package repositories

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}

func AddUser(ctx context.Context, user User) error {
	query :=
		`INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)`

	_, err := DB.Exec(ctx, query, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("error while inserting user: %w", err)
	}
	return nil
}
