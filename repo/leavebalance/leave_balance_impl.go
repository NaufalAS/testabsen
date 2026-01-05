package leavebalancerepo

import "test/model/domain"

type LeaveBalanceRepository interface {
	CreateLeaveBalance(lb domain.LeaveBalanve) (domain.LeaveBalanve, error)
	GetByUserAndType(userID int, leaveTypeID int) (domain.LeaveBalanve, error)
	DeductLeave(userId int,leaveTypeId int,year int,days int) error
}