package domain

import "time"

type LeaveApprovalLog struct {
	ID             int       `gorm:"column:id;primaryKey;autoIncrement"`
	LeaveRequestId int       `gorm:"column:leave_request_id"` 
	ApproverId     int       `gorm:"column:approver_id"`     
	Status         string    `gorm:"column:status"`           
	Comment        string    `gorm:"column:comment"`          
	CreatedAt      time.Time `gorm:"column:created_at"`       
}


func (LeaveApprovalLog) TableName() string {
	return "leave_approval_logs"
}