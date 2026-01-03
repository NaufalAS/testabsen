package domain

import "time"

// LeaveRequest mewakili tabel leave_requests
type LeaveRequest struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserId      int       `gorm:"column:user_id"`
	LeaveTypeId int       `gorm:"column:leave_type_id"`
	StartDate   time.Time `gorm:"column:start_date"`  // gunakan time.Time untuk tanggal
	EndDate     time.Time `gorm:"column:end_date"`
	Reason      string    `gorm:"column:reason"`
	Status      string    `gorm:"column:status"`      // bisa buat type custom leave_status jika mau
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// TableName override default table name di GORM
func (LeaveRequest) TableName() string {
	return "leave_requests"
}
