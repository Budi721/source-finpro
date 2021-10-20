package service

import (
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/model"
	"github.com/itp-backend/backend-a-co-create/repository"
)

type UserService interface {
	GetAllUser() []model.User
	FindByID(userID uint64) model.User
	UpdateUser(user dto.UserUpdateDTO) model.User
	ChangePassword(userID uint64, user dto.ChangePasswordDTO) model.User
	FindByEmail(email string) model.User
}

func GetAllUser() []model.User {
	res := repository.GetAllUser()
	return res
}

func GetAllJustUser() []model.User {
	res := repository.GetAllJustUser()
	return res
}

func FindByID(userID uint64) model.User {
	res := repository.Profile(userID)
	return res
}

func UpdateUser(userID uint64, user dto.UserUpdateDTO) model.User {
	userToUpdate := model.User{
		Name:  user.Name,
		Email: user.Email,
	}
	res := repository.UpdateUser(userID, userToUpdate)
	return res
}

func ChangePassword(userID uint64, user dto.ChangePasswordDTO) model.User {
	userToChangePassword := model.User{
		Password: user.Password,
	}
	res := repository.ChangePassword(userID, userToChangePassword)
	return res
}

func FindByEmail(email string) model.User {
	res := repository.FindByEmail(email)
	return res
}
