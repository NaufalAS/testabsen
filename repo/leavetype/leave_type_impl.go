package leavetyperepo

import "test/model/domain"

type LeaveTypeRepository interface {
	Create(lt domain.LeaveType) (domain.LeaveType, error)
	GetByID(id int) (domain.LeaveType, error)
	GetAll() ([]domain.LeaveType, error)
}