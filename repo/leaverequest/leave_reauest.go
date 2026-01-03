package leaverequestrepo

import (
	"test/model/domain"
	"time"

	"gorm.io/gorm"
)



type LeaveRequestRepositoryImpl struct {
	db *gorm.DB
}


func NewLeaveRequestRepository(db *gorm.DB) *LeaveRequestRepositoryImpl {
	return &LeaveRequestRepositoryImpl{db: db}
}

func (r *LeaveRequestRepositoryImpl) Create(request domain.LeaveRequest) (domain.LeaveRequest, error) {
	if err := r.db.Create(&request).Error; err != nil {
		return domain.LeaveRequest{}, err
	}
	return request, nil
}


func (r *LeaveRequestRepositoryImpl) GetByUserId(userId int) ([]domain.LeaveRequest, error) {
	var requests []domain.LeaveRequest
	if err := r.db.Where("user_id = ?", userId).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}


func (r *LeaveRequestRepositoryImpl) GetById(id int) (domain.LeaveRequest, error) {
	var request domain.LeaveRequest
	if err := r.db.First(&request, id).Error; err != nil {
		return domain.LeaveRequest{}, err
	}
	return request, nil
}


func (r *LeaveRequestRepositoryImpl) GetAll() ([]domain.LeaveRequest, error) {
	var requests []domain.LeaveRequest
	if err := r.db.Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *LeaveRequestRepositoryImpl) UpdateDates(id int, startDate, endDate time.Time) error {
	return r.db.Model(&domain.LeaveRequest{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"start_date": startDate,
			"end_date":   endDate,
			"updated_at": time.Now(),
		}).Error
}

func (r *LeaveRequestRepositoryImpl) UpdateStatus(id int, status string) error {
	return r.db.Model(&domain.LeaveRequest{}).
		Where("id = ?", id).
		Update("status", status).Error
}