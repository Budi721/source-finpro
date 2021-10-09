package model

type Role struct {
	GormModel
	Role   string `gorm:"not null;unique" json:"role"`
	UserID []User `json:"-"`
}
