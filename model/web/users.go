package userweb

type RegisterUserequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	JenisKelamin string `validate:"required" json:"jenis_kelamin"`
	NoTelephone string `validate:"required" json:"no_telephone"`
	RoleId 	 int `validate:"required" json:"role_id"`
}

type LoginUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UpdateUserRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	JenisKelamin string `validate:"required" json:"jenis_kelamin"`
	NoTelephone string `validate:"required" json:"no_telephone"`
	RoleId 	 int `validate:"reqiured" json:"role_id"`
}