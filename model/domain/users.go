package domain

import "time"


type Users struct{
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	EmployeeCode  string `gorm:"column:employee_code"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	JenisKelamin string `gorm:"column:jenis_kelamin"`
	NoTelephone string `gorm:"column:no_telephone"` 
	RoleId  int `gorm:"column:role_id"`
	CreatedAt time.Time
	UpdateAt time.Time
}

func (Users) TableName()  string  {
	return "users"
}