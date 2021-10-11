package dto

type UserUpdateDTO struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email" validate:"email"`
}

type ChangePasswordDTO struct {
	Password        string `json:"password" form:"password" binding:"required" validate:"min:6"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required" validate:"min:6"`
}
