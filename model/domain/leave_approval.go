package domain

import "time"

type LeaveApprovalLog struct {
	ID             int       `gorm:"column:id;primaryKey;autoIncrement"`
	LeaveRequestId int       `gorm:"column:leave_request_id"` // FK ke leave_requests.id
	ApproverId     int       `gorm:"column:approver_id"`      // FK ke users.id
	Status         string    `gorm:"column:status"`           // misal "pending", "approved", "rejected"
	Comment        string    `gorm:"column:comment"`          // optional komentar approver
	CreatedAt      time.Time `gorm:"column:created_at"`       // timestamp log dibuat
}

// TableName override default table name GORM
func (LeaveApprovalLog) TableName() string {
	return "leave_approval_logs"
}