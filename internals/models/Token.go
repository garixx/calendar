package models

type Token struct {
	Token string `json:"token" validate:"jwt"`
}

//go:generate mockgen -package mocks -destination ../mocks/token_usecase_mock.go . TokenUsecase
type TokenUsecase interface {
	CreateToken(payload Token) (Token, error)
	GetToken(token Token) (Token, error)
}

//go:generate mockgen -package mocks -destination ../mocks/token_repository_mock.go . TokenRepository
type TokenRepository interface {
	CreateToken(payload Token) (Token, error)
	GetToken(token Token) (Token, error)
}
