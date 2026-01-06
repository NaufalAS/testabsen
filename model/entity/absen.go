package entity

import (
	"test/model/domain"
	"time"
)

type LeaveApprovalEntity struct {
	ID           int       `json:"id"`
	ApproverID   int       `json:"approver_id"`
	ApproverName string    `json:"approver_name"`
	Status       string    `json:"status"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
}

type LeaveRequestEntity struct {
	ID          int                  `json:"id"`
	UserID      int                  `json:"user_id"`
	LeaveTypeID int                  `json:"leave_type_id"`
	StartDate   time.Time            `json:"start_date"`
	EndDate     time.Time            `json:"end_date"`
	Reason      string               `json:"reason"`
	Status      string               `json:"status"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Approvals   []LeaveApprovalEntity `json:"approvals"`
}


func ToLeaveApprovalEntity(log domain.LeaveApprovalLog, approverName string) LeaveApprovalEntity {
	return LeaveApprovalEntity{
		ID:           log.ID,
		ApproverID:   log.ApproverId,
		ApproverName: approverName,
		Status:       log.Status,
		Comment:      log.Comment,
		CreatedAt:    log.CreatedAt,
	}
}


func ToLeaveRequestEntity(leave domain.LeaveRequest, approvalLogs []domain.LeaveApprovalLog, approvers []domain.Users) LeaveRequestEntity {
	var logs []LeaveApprovalEntity
	for _, log := range approvalLogs {
		// ambil nama approver
		name := ""
		for _, u := range approvers {
			if u.ID == log.ApproverId {
				name = u.Name
				break
			}
		}
		logs = append(logs, ToLeaveApprovalEntity(log, name))
	}

	return LeaveRequestEntity{
		ID:          leave.ID,
		UserID:      leave.UserId,
		LeaveTypeID: leave.LeaveTypeId,
		StartDate:   leave.StartDate,
		EndDate:     leave.EndDate,
		Reason:      leave.Reason,
		Status:      leave.Status,
		CreatedAt:   leave.CreatedAt,
		UpdatedAt:   leave.UpdatedAt,
		Approvals:   logs,
	}
}
