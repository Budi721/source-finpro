package service

import (
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/model"
	"github.com/itp-backend/backend-a-co-create/repository"
)

type RoleService interface {
	GetAllRole() []model.Role
	FindRoleID(roleID uint64) model.Role
	CreateRole(role dto.RoleDTO) model.Role
	DeleteRole(roleID uint64) model.Role
}

func GetAllRole() []model.Role {
	res := repository.GetAllRole()
	return res
}

func FindRoleID(roleID uint64) model.Role {
	res := repository.FindRoleID(roleID)
	return res
}

func FindRoleForInject() model.Role {
	res := repository.FindRoleForInject()
	return res
}

func CreateRole(role dto.RoleDTO) model.Role {
	roleToCreate := model.Role{
		Role: role.Role,
	}
	res := repository.InsertRole(roleToCreate)
	return res
}

func CreateRoles(roles []model.Role) []model.Role {
	res := repository.InsertRoles(roles)
	return res
}

func DeleteRole(roleID uint64) model.Role {
	res := repository.DeleteRole(roleID)
	return res
}
