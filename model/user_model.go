package model

type User struct {
	GormModel
	Name     string `gorm:"not null" json:"name" form:"name"`
	Email    string `gorm:"not null;unique" json:"email" form:"email"`
	Password string `gorm:"->;<-;not null" json:"-" form:"password"`
	Token    string `gorm:"-" json:"token,omitempty"`
	RoleID   uint   `gorm:"->;<-;not null" json:"-"`
	Role     string `gorm:"-" json:"role"`
}
