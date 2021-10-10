package repository

import (
	"github.com/itp-backend/backend-a-co-create/config/database"
	"github.com/itp-backend/backend-a-co-create/helper/bc"
	"github.com/itp-backend/backend-a-co-create/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	InsertUser(user model.User) model.User
	// UpdateUser(userID string) model.User
	// ChangePassword(userID string) model.User
	VerifyCredential(email, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	GetAllUser() []model.User
	// Profile(userID string) model.User
}

var (
	db *gorm.DB = database.SetupDBConn()
)

func InsertUser(user model.User) model.User {
	user.Password = bc.HashAndSalt(user.Password)
	db.Save(&user)
	return user
}

// func UpdateUser(userID string) model.User {

// }

// func ChangePassword(userID string) model.User {

// }

func VerifyCredential(email, password string) interface{} {
	var user model.User
	err := db.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil
	}
	return user
}

func IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user model.User
	return db.Where("email = ?", email).Take(&user)
}

func GetAllUser() []model.User {
	var users []model.User
	db.Find(&users)
	return users
}

// func Profile(userID string) model.User {

// }
