package models

import (
	"time"
)

type UserRequest struct {
	Login    string `json:"login"    validate:"required,alphanum,tenMax"`
	Password string `json:"password" validate:"required,alphanum,twentyMax"`
}

type UserResponse struct {
	Id        string `json:"id"        validate:"required,uuid4"`
	Login     string `json:"login"     validate:"required,alphanum,tenMax"`
	CreatedAt string `json:"createdAt" validate:"required"`
}

type User struct {
	//Id           uuid.UUID `json:"id"        db:"id"`
	Id           string    `json:"id"        db:"id"            validate:"required,uuid4"`
	Login        string    `json:"login"     db:"login"         validate:"required,alphanum,tenMax"`
	PasswordHash string    `json:"~"         db:"password_hash" validate:"required"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	IsDeleted    bool      `json:"isDeleted" db:"is_deleted"`
}

//go:generate mockgen -package mocks -destination ../mocks/user_usecase_mock.go . UserUsecase
type UserUsecase interface {
	CreateUser(payload UserRequest) (User, error)
	GetUser(user User) (User, error)
}

//go:generate mockgen -package mocks -destination ../mocks/user_repository_mock.go . UserRepository
type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUser(user User) (User, error)
}
