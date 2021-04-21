package services

import (
	"example.com/mvc/domain"
	"example.com/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
