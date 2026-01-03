package leavetyperepo

import (
	"errors"
	"test/model/domain"

	"gorm.io/gorm"
)


type LeaveTypeRepositoryImpl struct {
	DB *gorm.DB
}

func NewLeaveTypeRepository(db *gorm.DB) *LeaveTypeRepositoryImpl {
	return &LeaveTypeRepositoryImpl{DB: db}
}

func (repo *LeaveTypeRepositoryImpl) Create(lt domain.LeaveType) (domain.LeaveType, error) {
	if err := repo.DB.Create(&lt).Error; err != nil {
		return domain.LeaveType{}, err
	}
	return lt, nil
}

func (repo *LeaveTypeRepositoryImpl) GetByID(id int) (domain.LeaveType, error) {
	var lt domain.LeaveType
	if err := repo.DB.Where("id = ?", id).Take(&lt).Error; err != nil {
		return domain.LeaveType{}, errors.New("leave type not found")
	}
	return lt, nil
}

func (repo *LeaveTypeRepositoryImpl) GetAll() ([]domain.LeaveType, error) {
	var leaveTypes []domain.LeaveType
	if err := repo.DB.Find(&leaveTypes).Error; err != nil {
		return nil, err
	}
	return leaveTypes, nil
}
