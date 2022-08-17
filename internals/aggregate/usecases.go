package aggregate

import (
	"calendar/internals/models"
	"log"
)

type Calendar struct {
	Logger   log.Logger
	UserCase models.UserUsecase
}

func NewCalendar(userCase models.UserUsecase) Calendar {
	return Calendar{
		UserCase: userCase,
	}
}
