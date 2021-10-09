package dto

// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
// 	RoleID   uint64 `json:"role_id,omitempty" form:"role_id,omitempty"`
// }

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
	RoleID   uint64 `json:"role_id,omitempty" form:"role_id,omitempty"`
}
