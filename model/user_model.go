package model

// type User struct {
// 	Id         int        `json:"id_user,omitempty"`
// 	Username   string     `json:"username,omitempty"`
// 	Password   string     `json:"password,omitempty"`
// 	LoginAs    string     `json:"login_as,omitempty"`
// 	AuthToken  string     `json:"-" gorm:"-"`
// 	Article    Article    `gorm:"foreignKey:IdUser" json:"-"`
// 	Project    Project    `gorm:"foreignKey:Admin" json:"-"`
// 	Enrollment Enrollment `gorm:"foreignKey:IdUser" json:"-"`
// }

type User struct {
	GormModel
	Name     string `gorm:"not null" json:"name" form:"name"`
	Email    string `gorm:"not null;unique" json:"email" form:"email"`
	Password string `gorm:"->;<-;not null" json:"-" form:"password"`
	Token    string `gorm:"-" json:"token,omitempty"`
	RoleID   uint   `gorm:"->;<-;not null" json:"-"`
	Role     Role   `gorm:"foreignKey:RoleID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"role,omitempty"`
}

// func (user *User) SetPassword(password string) string {
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
// 	return string(hashedPassword)
// }

// func (user *User) ComparePassword(password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// }

// type Admin User

/**
Logic Admin to verify
*/
