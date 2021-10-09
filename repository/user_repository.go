package repository

import (
	"github.com/itp-backend/backend-a-co-create/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	InsertUser(user model.User) model.User
	UpdateUser(user model.User) model.User
	VerifyCredential(email, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	GetAllUser(user model.User) []model.User
	Profile(userID string) model.User
}

func InsertUser(user model.User) model.User {

}

func UpdateUser(user model.User) model.User {

}

func VerifyCredential(email, password string) interface{} {

}

func IsDuplicateEmail(email string) (tx *gorm.DB) {

}

func GetAllUser(user model.User) []model.User {

}

func Profile(userID string) model.User {

}
