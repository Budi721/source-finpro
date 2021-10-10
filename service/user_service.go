package service

import (
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/model"
)

type UserService interface {
	GetAllUser() []model.User
	FindByEmail(email string) model.User
	UpdateUser(user dto.UserUpdateDTO) model.User
	ChangePassword(email, password string) model.User
}

// func GetAllUser() []model.User {

// }

// func FindByEmail(email string) model.User {

// }

// func UpdateUser(user dto.UserUpdateDTO) model.User {

// }

// func ChangePassword(email, password string) model.User {

// }
