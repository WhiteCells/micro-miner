package dao

import (
	"micro-user/dao/model"
	"micro-user/utils"
)

type UserDao struct {
}

func (UserDao) AddUser(user *model.User) error {
	return utils.DB.
		Create(user).Error
}

func (UserDao) DelUser(id int) error {
	return utils.DB.
		Delete(&model.User{}, id).Error
}

func (UserDao) UpdateUser(id int, user *model.User) error {
	return utils.DB.
		Save(user).Error
}

func (UserDao) UpdateUserPassword(id int, newPassword string) error {
	return utils.DB.
		Model(&model.User{}).
		Where("id=?", id).
		Update("password", newPassword).Error
}

func (UserDao) UpdateUserPoints(id int, points float32) error {
	return utils.DB.
		Model(&model.User{}).
		Where("id=?", id).
		Update("points", points).Error
}

func (UserDao) GetAllUsers() (*[]model.User, error) {
	var users []model.User
	err := utils.DB.
		Find(&users).Error
	return &users, err
}

func (UserDao) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := utils.DB.
		First(&user, id).Error
	return &user, err
}

func (UserDao) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := utils.DB.
		Where("email=?", email).
		First(&user).Error
	return &user, err
}

func (UserDao) GetUserByInviteCode(inviteCode string) (*model.User, error) {
	var user model.User
	err := utils.DB.
		Where("invite_code=?", inviteCode).
		First(&user).Error
	return &user, err
}
