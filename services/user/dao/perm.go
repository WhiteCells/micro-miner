package dao

import (
	"micro-user/model"
	"micro-user/utils"
)

type PermDao struct {
}

func (m *PermDao) AddPerm(perm model.Perm) error {
	return utils.DB.
		Create(perm).Error
}

func (m *PermDao) DelPerm(id int) error {
	return utils.DB.
		Delete(&model.Perm{}, id).Error
}

// func (m *PermDao) UpdatePerm(id int, )

func (m *PermDao) GetAllPerms() (*[]model.Perm, error) {
	var perms []model.Perm
	err := utils.DB.
		Find(&perms).Error
	return &perms, err
}

func (m *PermDao) AssignPermToRole(roleID int, permID int) error {
	var role model.Role
	var perm model.Perm

	if err := utils.DB.First(&role, roleID).Error; err != nil {
		return err
	}
	if err := utils.DB.First(&perm, permID).Error; err != nil {
		return err
	}

	return utils.DB.
		Model(&role).
		Association("perm").
		Append(&perm)
}
