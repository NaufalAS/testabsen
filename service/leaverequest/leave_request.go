package leaverequestservice

import (
	"errors"
	"test/model/domain"
	"test/model/entity"
	"time"

	leaveapprovalrepo "test/repo/leave_approval.go"
	leaverequestrepo "test/repo/leaverequest"
	userrepository "test/repo/user"
)

type LeaveApprovalService struct {
	approvalRepo leaveapprovalrepo.LeaveApprovalRepository
	requestRepo  leaverequestrepo.LeaveRequestRepository
	userRepo     userrepository.UserRepository
}


func NewLeaveApprovalService(ar leaveapprovalrepo.LeaveApprovalRepository, rr leaverequestrepo.LeaveRequestRepository, userRepo     userrepository.UserRepository) *LeaveApprovalService {
	return &LeaveApprovalService{
		approvalRepo: ar,
		requestRepo:  rr,
		userRepo: userRepo,
	}
}


func (s *LeaveApprovalService) CreateLeaveRequest(userId int, leaveTypeId int, startDate, endDate time.Time, reason string) (*domain.LeaveRequest, error) {

	newRequest := domain.LeaveRequest{
		UserId:      userId,
		LeaveTypeId: leaveTypeId,
		StartDate:   startDate,
		EndDate:     endDate,
		Reason:      reason,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	request, err := s.requestRepo.Create(newRequest)
	if err != nil {
		return nil, err
	}

	approvers, err := s.userRepo.GetUsersByRoleIds([]int{2, 3})
	if err != nil {
		return nil, err
	}

	for _, approver := range approvers {
	log := domain.LeaveApprovalLog{
		LeaveRequestId: request.ID,
		ApproverId:     approver.ID, 
		Status:         "pending",
		CreatedAt:      time.Now(),
	}
	_, err := s.approvalRepo.Create(log)
	if err != nil {
		return nil, err
	}
}


	return &request, nil
}

func (s *LeaveApprovalService) GetLeaveRequests() ([]entity.LeaveRequestEntity, error) {
	
	leaves, err := s.requestRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []entity.LeaveRequestEntity

	for _, leave := range leaves {
		
		logs, err := s.approvalRepo.GetByLeaveRequestId(leave.ID)
		if err != nil {
			return nil, err
		}

		
		var approverIDs []int
		for _, log := range logs {
			approverIDs = append(approverIDs, log.ApproverId)
		}
		approvers, err := s.userRepo.GetUsersByIDs(approverIDs)
		if err != nil {
			return nil, err
		}

		
		result = append(result, entity.ToLeaveRequestEntity(leave, logs, approvers))
	}

	return result, nil
}

func (s *LeaveApprovalService) UpdateLeaveDates(leaveId int, startDate, endDate time.Time) error {
	_, err := s.requestRepo.GetById(leaveId)
	if err != nil {
		return err
	}

	
	logs, _ := s.approvalRepo.GetByLeaveRequestId(leaveId)
	for _, l := range logs {
		if l.Status != "pending" {
			return errors.New("Tidak bisa edit, salah satu approval sudah selesai")
		}
	}

	
	return s.requestRepo.UpdateDates(leaveId, startDate, endDate)
}

func (s *LeaveApprovalService) ApproveLeave(leaveId, approverId int, status, comment string) error {
	logs, err := s.approvalRepo.GetByLeaveRequestId(leaveId)
	if err != nil {
		return err
	}

	var current *domain.LeaveApprovalLog
	for i := range logs {
		if logs[i].Status == "pending" {
			current = &logs[i]
			break
		}
	}

	if current == nil {
		return errors.New("Semua approval sudah selesai")
	}

	if current.ApproverId != approverId {
		return errors.New("Bukan giliran Anda untuk approve")
	}

	
	current.Status = status
	current.Comment = comment
	if err := s.approvalRepo.UpdateStatus(current.ID, status, comment); err != nil {
	return err
}

	
	leave, err := s.requestRepo.GetById(leaveId)
	if err != nil {
		return err
	}

	if status == "rejected" {
		return s.requestRepo.UpdateStatus(leave.ID, "rejected")
	} else {
		
		allApproved := true
		for _, l := range logs {
			if l.Status != "approved" && l.ID != current.ID {
				allApproved = false
				break
			}
		}
		if allApproved {
			return s.requestRepo.UpdateStatus(leave.ID, "approved")
		}
	}

	return nil
}
