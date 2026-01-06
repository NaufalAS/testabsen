package domain

import "time"

// LeaveRequest mewakili tabel leave_requests
type LeaveRequest struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserId      int       `gorm:"column:user_id"`
	LeaveTypeId int       `gorm:"column:leave_type_id"`
	StartDate   time.Time `gorm:"column:start_date"`  
	EndDate     time.Time `gorm:"column:end_date"`
	Reason      string    `gorm:"column:reason"`
	Status      string    `gorm:"column:status"`      
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}


func (LeaveRequest) TableName() string {
	return "leave_requests"
}
