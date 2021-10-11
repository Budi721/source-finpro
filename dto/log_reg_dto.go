package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
	RoleID   uint64 `json:"role_id" form:"role_id" binding:"required"`
}

type ResponseLogRegDTO struct {
	ID    uint64 `json:"id" form:"id" binding:"required"`
	Name  string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Token string `json:"token" form:"token" binding:"required" validate:"token"`
}
