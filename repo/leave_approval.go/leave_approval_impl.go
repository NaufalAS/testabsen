package leaveapprovalrepo

import "test/model/domain"

type LeaveApprovalRepository interface {
	Create(log domain.LeaveApprovalLog) (domain.LeaveApprovalLog, error)
	GetByLeaveRequestId(leaveRequestId int) ([]domain.LeaveApprovalLog, error)
	GetById(id int) (domain.LeaveApprovalLog, error)
	UpdateStatus(id int, status string, comment string) error
}