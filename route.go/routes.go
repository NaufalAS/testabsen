package routego

import (
	"os"
	"test/app"
	leaverequestcontroller "test/controller/leaverequest"
	usercontroller "test/controller/user"
	"test/helper"
	leaveapprovalrepo "test/repo/leave_approval.go"
	leavebalancerepo "test/repo/leavebalance"
	leaverequestrepo "test/repo/leaverequest"
	leavetyperepo "test/repo/leavetype"
	userrepository "test/repo/user"
	leaverequestservice "test/service/leaverequest"
	userservice "test/service/user"

	"github.com/labstack/echo/v4"
)

func UserRoutes (prefix string, e *echo.Echo) {
	db := app.Dbconncentio()
	

	leaveBalaancerepo := leavebalancerepo.NewLeaveBalanceRepository(db)
	leavetyperepo := leavetyperepo.NewLeaveTypeRepository(db)
	userAuthrepo := userrepository.NewUserRepository(db)
	userAuthService := userservice.NewUserServic(userAuthrepo, leaveBalaancerepo, leavetyperepo)
	userAuthContller := usercontroller.NewAuthController(userAuthService)

	leaveapprovalrepo := leaveapprovalrepo.NewLeaveApprovalRepository(db)
	leaverequestrepo := leaverequestrepo.NewLeaveRequestRepository(db)
	leavereuestservice := leaverequestservice.NewLeaveApprovalService(leaveapprovalrepo,leaverequestrepo, userAuthrepo, leaveBalaancerepo)
	leaverequuestcontroller := leaverequestcontroller.NewLeaveController(leavereuestservice)



	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", userAuthContller.Register)
	authRoute.POST("/login", userAuthContller.Login)

	leaveroute := g.Group("/leave")
	leaveroute.Use(helper.JWTMiddleware(os.Getenv("SECRET_KEY"))) 
	leaveroute.POST("/create", leaverequuestcontroller.CreateLeaveRequest)
	leaveroute.GET("/list", leaverequuestcontroller.GetLeaveRequests)
	leaveroute.PUT("/update-dates", leaverequuestcontroller.UpdateLeaveDates)        // update start/end date (hanya pending)
	leaveroute.PUT("/approve", leaverequuestcontroller.ApproveLeave)   
}