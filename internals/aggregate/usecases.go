package aggregate

import (
	"calendar/internals/models"
	"log"
)

type Calendar struct {
	Logger    log.Logger
	UserCase  models.UserUsecase
	TokenCase models.TokenUsecase
}

//func NewCalendar(userCase models.UserUsecase) Calendar {
//	return Calendar{
//		UserCase: userCase,
//	}
//}
