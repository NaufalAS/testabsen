package entity

import "test/model/domain"

type UserEntity struct {
	UserID int    `json:"id"`
	EmployeeCode string `json:"employee_code"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoTelephone string `json:"no_telephone"`
	RoleId int `json:"role_id"`
}

func ToUserEntity(user domain.Users) UserEntity {
	return UserEntity{
		UserID: user.ID,
		EmployeeCode: user.EmployeeCode,
		Name: user.Name,
		Email: user.Email,
		JenisKelamin: user.JenisKelamin,
		NoTelephone: user.NoTelephone,
		RoleId: user.RoleId,
	}
}

func ToUserListEntity(users []domain.Users) []UserEntity {
	var userData []UserEntity

	for _, user := range users {
		userData = append(userData, ToUserEntity(user))

	}
	return userData
}