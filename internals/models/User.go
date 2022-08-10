package models

import (
	"github.com/google/uuid"
	"time"
)

type SignUpPayload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	Id           uuid.UUID `json:"id"        db:"id"`
	Login        string    `json:"login"     db:"login"`
	PasswordHash string    `json:"~"         db:"password_hash"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	IsDeleted    bool      `json:"isDeleted" db:"is_deleted"`
}

type UserUsecase interface {
	CreateUser(payload SignUpPayload) (User, error)
	GetUser(user User) (User, error)
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUser(user User) (User, error)
}
