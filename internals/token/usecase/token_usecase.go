package usecase

import "calendar/internals/models"

type TokenUsecase struct {
	tokenRepo models.TokenRepository
}

func NewTokenUsecase(repo models.TokenRepository) models.TokenUsecase {
	return &TokenUsecase{
		tokenRepo: repo,
	}
}

func (t *TokenUsecase) CreateToken(payload models.Token) (models.Token, error) {
	return t.tokenRepo.CreateToken(payload)
}

func (t *TokenUsecase) GetToken(token models.Token) (models.Token, error) {
	return t.tokenRepo.GetToken(token)
}
