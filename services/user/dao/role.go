package dao

import (
	"micro-user/model"
	"micro-user/utils"
)

type RoleDao struct {
}

func (RoleDao) AddRole(role model.Role) error {
	return utils.DB.
		Create(role).Error
}

func (RoleDao) DelRole(role model.Role)
