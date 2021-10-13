package repository

import (
    "github.com/itp-backend/backend-a-co-create/config/database"
    "github.com/itp-backend/backend-a-co-create/helper/bc"
    "github.com/itp-backend/backend-a-co-create/model"
    "gorm.io/gorm"
)

type UserRepo interface {
	InsertUser(user model.User) model.User
	VerifyCredential(email, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	GetAllUser() []model.User
	UpdateUser(user model.User) model.User
	ChangePassword(userID uint64, user model.User) model.User
	Profile(userID uint64) model.User
	FindByEmail(email string) model.User
}

var (
	db *gorm.DB = database.SetupDBConn()
)

func InsertUser(user model.User, enrollment model.Enrollment) model.Enrollment {
	tx := db.Begin()

	user.Password = bc.HashAndSalt(user.Password)
	result := tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return model.Enrollment{}
	}

    enrollment.IdUser = user.ID
    enrollment.Password = user.Password
	result = tx.Create(&enrollment)
	if result.Error != nil {
		tx.Rollback()
		return model.Enrollment{}
    }

    tx.Commit()
	return enrollment
}

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

func UpdateUser(userID uint64, user model.User) model.User {
	var tempUser model.User
	db.First(&tempUser, userID)
	user.ID = tempUser.ID
	user.Password = tempUser.Password
	user.RoleID = tempUser.RoleID

	db.Save(&user)
	return user
}

func ChangePassword(userID uint64, user model.User) model.User {
	var tempUser model.User
	db.First(&tempUser, userID)
	user.ID = tempUser.ID
	user.Name = tempUser.Name
	user.Email = tempUser.Email
	user.RoleID = tempUser.RoleID

	user.Password = bc.HashAndSalt(user.Password)
	db.Save(&user)
	return user
}

func Profile(userID uint64) model.User {
	var user model.User
	db.First(&user, userID)
	return user
}

func FindByEmail(email string) model.User {
	var user model.User
	db.Where("email = ?", email).Take(&user)
	return user
}
