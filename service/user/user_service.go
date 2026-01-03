package userservice

import (
	"errors"
	"test/helper"
	"test/model/domain"
	userweb "test/model/web"
	leavebalancerepo "test/repo/leavebalance"
	leavetyperepo "test/repo/leavetype"
	userrepository "test/repo/user"
	leavebalance "test/service/leave_balance"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepo     userrepository.UserRepository
	lbRepo   leavebalancerepo.LeaveBalanceRepository // <- tambahkan ini
	ltRepo   leavetyperepo.LeaveTypeRepository 
}

func NewUserServic(userRepo    userrepository.UserRepository, lbRepo   leavebalancerepo.LeaveBalanceRepository ,ltRepo   leavetyperepo.LeaveTypeRepository ) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo:     userRepo,
		lbRepo: lbRepo,
		ltRepo: ltRepo,
	}
}

func (service *UserServiceImpl) Register(req userweb.RegisterUserequest) (helper.CustomResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	newUser := domain.Users{
		EmployeeCode: helper.GenerateEmployeeCode(),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passHash),
		JenisKelamin: req.JenisKelamin,
		NoTelephone: req.NoTelephone,
		RoleId: req.RoleId,
	}

	result, err := service.userRepo.Register(newUser)

	if err != nil {
		return nil, err
	}
leaveService := leavebalance.NewLeaveBalanceService(service.lbRepo, service.ltRepo)
if err := leaveService.InitLeaveBalance(result); err != nil {
	return nil, err
}

	return helper.CustomResponse{
		"name":  result.Name,
		"email": result.Email,
	}, nil
}


func (service *UserServiceImpl) Login(req userweb.LoginUserRequest) (helper.CustomResponse, error) {
	user, getuserErr := service.userRepo.Login(req.Email)
	if getuserErr != nil {
		return nil, errors.New("wrong email or password")
	}

	if checkPasswordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); checkPasswordErr != nil {
		return nil, errors.New("wrong email or password")
	}

	loginResponse, loginErr := helper.Login(user.ID, user.RoleId, user.Name, user.Email)
	if loginErr != nil {
		return nil, loginErr
	}

	return helper.CustomResponse{
		"token":      loginResponse.Token,
		"expires_at": loginResponse.ExpiresAt,
	}, nil
}
