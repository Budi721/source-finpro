package repository

import "github.com/itp-backend/backend-a-co-create/model"

type RoleRepo interface {
	GetAllRole() []model.Role
	InsertRole(role model.Role) model.Role
	FindRoleID(roleID uint64) model.Role
	DeleteRole(roleID uint64) model.Role
}

func GetAllRole() []model.Role {
	var roles []model.Role
	db.Find(&roles)
	return roles
}

func InsertRole(role model.Role) model.Role {
	db.Save(&role)
	return role
}

func InsertRoles(roles []model.Role) []model.Role {
	db.Create(&roles)
	return roles
}

func FindRoleID(roleID uint64) model.Role {
	var role model.Role
	db.First(&role, roleID)
	return role
}

func FindRoleForInject() model.Role {
	var role model.Role
	db.First(&role)
	return role
}

func DeleteRole(roleID uint64) model.Role {
	var role model.Role
	db.Delete(&role, roleID)
	return role
}
