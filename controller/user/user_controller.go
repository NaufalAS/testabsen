package usercontroller

import (
	"net/http"
	"test/helper"
	userweb "test/model/web"
	userservice "test/service/user"

	"github.com/labstack/echo/v4"
)

// UserControllerImpl is the implementation of AuthController interface
type UserControllerImpl struct {
	userservice userservice.Userservice
}

// NewAuthController creates a new instance of UserControllerImpl
func NewAuthController(userservice userservice.Userservice) *UserControllerImpl {
	return &UserControllerImpl{
		userservice: userservice,
	}
}

func (controller *UserControllerImpl) Register(c echo.Context) error {
	newUser := new(userweb.RegisterUserequest)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newUser); err != nil {
		return err
	}

	result, err := controller.userservice.Register(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Register Success", result))
}

func (controller *UserControllerImpl) Login(c echo.Context) error {
	user := new(userweb.LoginUserRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	result, err := controller.userservice.Login(*user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Login Success", result))
}