package leavebalancerepo

import (
	"errors"
	"test/model/domain"

	"gorm.io/gorm"
)

type LeaveBalanceRepositoryImpl struct {
	DB *gorm.DB
}

func NewLeaveBalanceRepository(db *gorm.DB) *LeaveBalanceRepositoryImpl {
	return &LeaveBalanceRepositoryImpl{DB: db}
}

func (repo *LeaveBalanceRepositoryImpl) CreateLeaveBalance(lb domain.LeaveBalanve) (domain.LeaveBalanve, error) {
	if err := repo.DB.Create(&lb).Error; err != nil {
		return domain.LeaveBalanve{}, err
	}
	return lb, nil
}

func (repo *LeaveBalanceRepositoryImpl) GetByUserAndType(userID int, leaveTypeID int) (domain.LeaveBalanve, error) {
	var lb domain.LeaveBalanve
	if err := repo.DB.Where("user_id = ? AND leave_type_id = ?", userID, leaveTypeID).Take(&lb).Error; err != nil {
		return domain.LeaveBalanve{}, errors.New("leave balance not found")
	}
	return lb, nil
}


func (repo *LeaveBalanceRepositoryImpl) DeductLeave(userId int,leaveTypeId int,year int,days int) error {
	return repo.DB.Model(&domain.LeaveBalanve{}).
		Where(
			"user_id = ? AND leave_type_id = ? AND year = ?",
			userId,
			leaveTypeId,
			year,
		).
		Updates(map[string]interface{}{
			"used_days":      gorm.Expr("used_days + ?", days),
			"remaining_days": gorm.Expr("remaining_days - ?", days),
		}).Error
}
