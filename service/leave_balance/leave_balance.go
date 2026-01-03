package leavebalance

import (
	"test/model/domain"
	leavebalancerepo "test/repo/leavebalance"
	leavetyperepo "test/repo/leavetype"
	"time"
)

type LeaveBalanceServiceImpl struct {
	lbRepo leavebalancerepo.LeaveBalanceRepository
	ltRepo leavetyperepo.LeaveTypeRepository
}

func NewLeaveBalanceService(lbRepo leavebalancerepo.LeaveBalanceRepository, ltRepo leavetyperepo.LeaveTypeRepository) *LeaveBalanceServiceImpl {
	return &LeaveBalanceServiceImpl{
		lbRepo: lbRepo,
		ltRepo: ltRepo,
	}
}

func (s *LeaveBalanceServiceImpl) InitLeaveBalance(user domain.Users) error {
	leaveTypes, err := s.ltRepo.GetAll()
	if err != nil {
		return err
	}

for _, lt := range leaveTypes {
	
	if user.JenisKelamin == "laki-laki" && lt.NameLeave == "Cuti Melahirkan" {
		continue
	}


		lb := domain.LeaveBalanve{
			UserId:      user.ID,
			LeaveTypeId: lt.ID,
			TotalDay:      lt.MaxDays, 
	UsedDay:       0,                
	RemainingDays: lt.MaxDays, 
	Year:          time.Now().Year(),
	UpdateAt:      time.Now(),
		}
		if _, err := s.lbRepo.CreateLeaveBalance(lb); err != nil {
			return err
		}
	}



	return nil
}