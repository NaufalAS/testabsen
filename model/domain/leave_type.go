package domain

import "time"

type LeaveType struct {
	ID                 int       `gorm:"column:id;primaryKey;autoIncrement"`
	NameLeave          string    `gorm:"column:names_leave"`         
	MaxDays            int       `gorm:"column:max_days"`            
	IsPaid             bool      `gorm:"column:is_paid"`             
	RequiresAttachment bool      `gorm:"column:requires_attachment"` 
	GenderOnly			string	 `gorm:"column:gender_only"`
	Description        string    `gorm:"column:description"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdateAt           time.Time `gorm:"column:update_at"`
}

func (LeaveType) TableName()  string  {
	return "leave_type"
}