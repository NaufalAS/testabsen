package domain

import "time"

type LeaveBalanve struct{
	ID int `gorm:"column:id;primarykey;auyoIncrement"`
	UserId int  `gorm:"column:user_id"`
	LeaveTypeId int `gorm:"column:leave_type_id"`
	TotalDay int `gorm:"column:total_days"`
	UsedDay int `gorm:"column:used_days"`
	RemainingDays int `gorm:"column:remaining_days"`
	Year int `gorm:"column:year"`
	UpdateAt time.Time `gorm:"column:updated_at"`
}

func (LeaveBalanve) TableName()  string  {
	return "leave_balances"
}