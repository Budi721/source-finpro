package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
	LoginAs  uint64 `json:"login_as" form:"login_as" binding:"required"`
}

type RegisterDTO struct {
	Name          string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email         string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password      string `json:"password" form:"password" binding:"required" validate:"min:6"`
	RoleID        uint64 `json:"role_id" form:"role_id" binding:"required"`
	TopikDiminati string `json:"topik_diminati,omitempty" binding:"required"`
}

type ResponseLogRegDTO struct {
	ID            uint64 `json:"id" form:"id" binding:"required"`
	Name          string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email         string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Role          string `json:"role" form:"role" binding:"required"`
	Token         string `json:"token" form:"token" binding:"required" validate:"token"`
	TopikDiminati string `json:"topik_diminati,omitempty"`
}
