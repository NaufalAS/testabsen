package userservice

import (
	"test/helper"
	userweb "test/model/web"
)

type Userservice interface{
	Register(req userweb.RegisterUserequest) (helper.CustomResponse, error)
	Login(req userweb.LoginUserRequest) (helper.CustomResponse, error)
}