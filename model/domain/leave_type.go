package domain

import "time"

type LeaveType struct {
	ID                 int       `gorm:"column:id;primaryKey;autoIncrement"`
	NameLeave          string    `gorm:"column:names_leave"`         // Nama cuti
	MaxDays            int       `gorm:"column:max_days"`            // Jumlah hari maksimal
	IsPaid             bool      `gorm:"column:is_paid"`             // Apakah cuti dibayar
	RequiresAttachment bool      `gorm:"column:requires_attachment"` // Perlu bukti atau lampiran
	GenderOnly			string	 `gorm:"column:gender_only"`
	Description        string    `gorm:"column:description"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdateAt           time.Time `gorm:"column:update_at"`
}

func (LeaveType) TableName()  string  {
	return "leave_type"
}