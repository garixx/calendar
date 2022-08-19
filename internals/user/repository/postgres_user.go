package repository

import (
	"calendar/internals/database"
	"calendar/internals/models"
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

type PostgresUserRepository struct {
	client database.Client
}

func NewPostgresUserRepository(client database.Client) *PostgresUserRepository {
	return &PostgresUserRepository{client: client}
}

func (p *PostgresUserRepository) CreateUser(user models.User) (models.User, error) {
	var newUser models.User
	query := "INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id, login, password_hash, created_at, is_deleted"
	row := p.client.QueryRow(context.Background(), query, user.Login, user.PasswordHash)

	if err := row.Scan(
		&newUser.Id,
		&newUser.Login,
		&newUser.PasswordHash,
		&newUser.CreatedAt,
		&newUser.IsDeleted,
	); err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (p *PostgresUserRepository) GetUserByLogin(login string) (models.User, error) {
	var users []*models.User
	err := pgxscan.Select(context.Background(), p.client, &users, "select * from users where login = $1 ", login)
	if err != nil {
		return models.User{}, err
	}
	if len(users) < 1 {
		return models.User{}, pgx.ErrNoRows
	}

	return *users[0], nil
}
