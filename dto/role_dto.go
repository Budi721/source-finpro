package dto

type RoleDTO struct {
	Role string `json:"role" form:"role" binding:"required"`
}
