package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id         int        `json:"id_user,omitempty"`
	Username   string     `json:"username,omitempty"`
	Password   string     `json:"password,omitempty"`
	LoginAs    string     `json:"login_as,omitempty"`
	AuthToken  string     `json:"-" gorm:"-"`
	Article    Article    `gorm:"foreignKey:IdUser" json:"-"`
	Project    Project    `gorm:"foreignKey:Admin" json:"-"`
	Enrollment Enrollment `gorm:"foreignKey:IdUser" json:"-"`
}

func (user *User) SetPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedPassword)
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

type Admin User

/**
Logic Admin to verify
*/
