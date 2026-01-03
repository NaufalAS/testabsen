package userweb





type LeaveRequestCreateRequest struct {
	Userid int `json:"user_id"`
	LeaveTypeId int    `json:"leave_type_id"`
	StartDate   string `json:"start_date"` // format YYYY-MM-DD
	EndDate     string `json:"end_date"`   // format YYYY-MM-DD
	Reason      string `json:"reason"`
}


type LeaveRequestUpdateDates struct {
	LeaveId   int       `json:"leave_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}



type LeaveRequestApproveRequest struct {
	LeaveId int    `json:"leave_id"`
	Status  string `json:"status"`  // "approved" atau "rejected"
	Comment string `json:"comment"`
}
