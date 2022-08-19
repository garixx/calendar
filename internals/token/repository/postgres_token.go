package repository

import (
	"calendar/internals/database"
	"calendar/internals/models"
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

type PostgresTokenRepository struct {
	client database.Client
}

func NewPostgresUserRepository(client database.Client) *PostgresTokenRepository {
	return &PostgresTokenRepository{client: client}
}

func (p *PostgresTokenRepository) CreateToken(payload models.Token) (models.Token, error) {
	var newToken models.Token
	query := "INSERT INTO tokens (token) VALUES ($1) RETURNING token"
	row := p.client.QueryRow(context.Background(), query, payload.Token)

	if err := row.Scan(&newToken.Token); err != nil {
		return models.Token{}, err
	}
	return newToken, nil
}

func (p *PostgresTokenRepository) GetToken(token models.Token) (models.Token, error) {
	var tokens []*models.Token
	err := pgxscan.Select(context.Background(), p.client, &tokens, "select token from tokens where token = $1", token.Token)
	if err != nil {
		return models.Token{}, err
	}
	if len(tokens) < 1 {
		return models.Token{}, pgx.ErrNoRows
	}

	return *tokens[0], nil
}
