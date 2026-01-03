package leaverequestrepo

import (
	"test/model/domain"
	"time"
)

type LeaveRequestRepository interface {
	Create(request domain.LeaveRequest) (domain.LeaveRequest, error)
	GetByUserId(userId int) ([]domain.LeaveRequest, error)
	GetById(id int) (domain.LeaveRequest, error)
	GetAll() ([]domain.LeaveRequest, error)
	UpdateDates(id int, startDate, endDate time.Time) error
	UpdateStatus(id int, status string) error
}