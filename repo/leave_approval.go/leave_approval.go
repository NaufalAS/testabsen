package leaveapprovalrepo

import (
	"test/model/domain"
	"gorm.io/gorm"
)



type LeaveApprovalRepositoryImpl struct {
	db *gorm.DB
}


func NewLeaveApprovalRepository(db *gorm.DB) *LeaveApprovalRepositoryImpl {
	return &LeaveApprovalRepositoryImpl{db: db}
}


func (r *LeaveApprovalRepositoryImpl) Create(log domain.LeaveApprovalLog) (domain.LeaveApprovalLog, error) {
	if err := r.db.Create(&log).Error; err != nil {
		return domain.LeaveApprovalLog{}, err
	}
	return log, nil
}


func (r *LeaveApprovalRepositoryImpl) GetByLeaveRequestId(leaveRequestId int) ([]domain.LeaveApprovalLog, error) {
	var logs []domain.LeaveApprovalLog
	if err := r.db.Where("leave_request_id = ?", leaveRequestId).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *LeaveApprovalRepositoryImpl) GetById(id int) (domain.LeaveApprovalLog, error) {
	var log domain.LeaveApprovalLog
	if err := r.db.First(&log, id).Error; err != nil {
		return domain.LeaveApprovalLog{}, err
	}
	return log, nil
}


func (r *LeaveApprovalRepositoryImpl) UpdateStatus(id int, status string, comment string) error {
	return r.db.Model(&domain.LeaveApprovalLog{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":  status,
			"comment": comment,
		}).Error
}

