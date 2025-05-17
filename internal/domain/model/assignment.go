package model

type AssignmentDetail struct {
	ShiftID   int    `json:"shift_id"`
	WorkerID  int    `json:"worker_id"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Role      string `json:"role"`
	Location  string `json:"location"`
}
