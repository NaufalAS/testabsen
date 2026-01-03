package userrepository

import (
	"errors"
	"test/model/domain"

	"gorm.io/gorm"
)


type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repo *UserRepositoryImpl) Register(user domain.Users) (domain.Users, error) {
	if err := repo.DB.Create(&user).Error; err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) Login(email string) (domain.Users, error) {
	var user domain.Users
	if err := repo.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		return domain.Users{}, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) UpdateUser(userID int, user domain.Users) (domain.Users, error) {
	if err := repo.DB.Model(&domain.Users{}).Where("id = ?", userID).Updates(user).Error; err != nil {
		return domain.Users{}, errors.New("failed to update user")
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetByID(userID int) (domain.Users, error) {
	var user domain.Users
	if err := repo.DB.Where("id = ?", userID).Take(&user).Error; err != nil {
		return domain.Users{}, errors.New("user not found")
	}
	return user, nil
}

func (repo *UserRepositoryImpl) GetUsersByRoleIds(roleIds []int) ([]domain.Users, error) {
	var users []domain.Users
	if err := repo.DB.Where("role_id IN ?", roleIds).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepositoryImpl) GetUsersByIDs(ids []int) ([]domain.Users, error) {
	var users []domain.Users
	if err := repo.DB.Where("id IN ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
