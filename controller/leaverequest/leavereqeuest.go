package leaverequestcontroller

import (
	"net/http"
	"test/helper"
	"test/model/web"
	leaverequestservice "test/service/leaverequest"
	"time"

	"github.com/labstack/echo/v4"
)

type LeaveController struct {
	leaveService *leaverequestservice.LeaveApprovalService
}

// Constructor
func NewLeaveController(ls *leaverequestservice.LeaveApprovalService) *LeaveController {
	return &LeaveController{
		leaveService: ls,
	}
}

// POST /api/leave-request
func (c *LeaveController) CreateLeaveRequest(ctx echo.Context) error {
	req := new(userweb.LeaveRequestCreateRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	// parse userId dari param atau token
	userId := helper.GetAuthId(ctx) 

	// parse tanggal
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}
	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	// panggil service
	leaveRequest, err := c.leaveService.CreateLeaveRequest(userId, req.LeaveTypeId, startDate, endDate, req.Reason)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Register Success", leaveRequest))
}


func (c *LeaveController) GetLeaveRequests(ctx echo.Context) error {
	leaves, err := c.leaveService.GetLeaveRequests()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", leaves))
}

//
// Update tanggal leave
func (c *LeaveController) UpdateLeaveDates(ctx echo.Context) error {
	req := new(userweb.LeaveRequestUpdateDates)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	// Parse tanggal dari string ke time.Time
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, "Format start_date salah", nil))
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, "Format end_date salah", nil))
	}

	// Panggil service
	if err := c.leaveService.UpdateLeaveDates(req.LeaveId, startDate, endDate); err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Tanggal leave berhasil diupdate", nil))
}



// Approve / Reject leave
func (c *LeaveController) ApproveLeave(ctx echo.Context) error {
	req := new(userweb.LeaveRequestApproveRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	userId := helper.GetAuthId(ctx)
	err := c.leaveService.ApproveLeave(req.LeaveId, userId, req.Status, req.Comment)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Approval berhasil", nil))
}