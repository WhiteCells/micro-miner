package dao

import (
	"micro-user/dao/model"
	"micro-user/utils"
)

type RoleDao struct {
}

func (RoleDao) AddRole(role model.Role) error {
	return utils.DB.
		Create(role).Error
}

func (RoleDao) DelRole(id int) error {
	return utils.DB.
		Delete(model.Role{}, id).Error
}

func (RoleDao) GetAllRoles() (*[]model.Role, error) {
	var roles []model.Role
	err := utils.DB.
		Find(&roles).Error
	return &roles, err
}

func (RoleDao) AssignRoleToUser(roleID, userID int) error {
	var role model.Role
	var user model.User
	if err := utils.DB.First(&role, roleID).Error; err != nil {
		return err
	}
	if err := utils.DB.First(&user, userID).Error; err != nil {
		return err
	}
	return utils.DB.
		Model(&role).
		Association("Users").
		Append(&user)
}
