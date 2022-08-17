package usecase

import (
	"calendar/internals/hashing"
	"calendar/internals/models"
)

type UserUsecase struct {
	userRepo models.UserRepository
}

func NewUserUsecase(repo models.UserRepository) models.UserUsecase {
	return &UserUsecase{
		userRepo: repo,
	}
}

func (u *UserUsecase) CreateUser(payload models.SignUpPayload) (models.User, error) {
	hash, err := hashing.HashPassword(payload.Password)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		Login:        payload.Login,
		PasswordHash: hash,
	}
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) GetUser(user models.User) (models.User, error) {
	return u.userRepo.GetUser(user)
}
