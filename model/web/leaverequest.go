package userweb





type LeaveRequestCreateRequest struct {
	Userid int `json:"user_id"`
	LeaveTypeId int    `json:"leave_type_id"`
	StartDate   string `json:"start_date"` 
	EndDate     string `json:"end_date"`   
	Reason      string `json:"reason"`
}


type LeaveRequestUpdateDates struct {
	LeaveId   int       `json:"leave_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}



type LeaveRequestApproveRequest struct {
	LeaveId int    `json:"leave_id"`
	Status  string `json:"status"`  
	Comment string `json:"comment"`
}
