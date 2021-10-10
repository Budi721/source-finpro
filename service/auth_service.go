package service

import (
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/helper/bc"
	"github.com/itp-backend/backend-a-co-create/model"
	"github.com/itp-backend/backend-a-co-create/repository"
)

type AuthService interface {
	VerifyCredential(email, password string) interface{}
	CreateUser(user dto.RegisterDTO) model.User
	IsDuplicateEmail(email string) bool
}

func VerifyCredential(email, password string) interface{} {
	res := repository.VerifyCredential(email, password)
	if val, ok := res.(model.User); ok {
		comparePass := bc.ComparePass(val.Password, password)
		if val.Email == email && comparePass {
			return res
		}
		return false
	}
	return false
}

func CreateUser(user dto.RegisterDTO) model.User {
	userToCreate := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   uint(user.RoleID),
	}
	res := repository.InsertUser(userToCreate)
	return res
}

func IsDuplicateEmail(email string) bool {
	err := repository.IsDuplicateEmail(email).Error
	return err == nil
}
