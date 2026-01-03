package userrepository

import "test/model/domain"

type UserRepository interface {
	Register(user domain.Users) (domain.Users, error) 
	Login(email string) (domain.Users, error)
	UpdateUser(userID int, user domain.Users) (domain.Users, error)
	GetByID(userID int) (domain.Users, error)
	GetUsersByRoleIds(roleIds []int) ([]domain.Users, error)
	GetUsersByIDs(ids []int) ([]domain.Users, error)
}