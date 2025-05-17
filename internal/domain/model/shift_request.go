package model

type ShiftRequest struct {
	ID        int    `json:"id"`
	ShiftID   int    `json:"shift_id"`
	WorkerID  int    `json:"worker_id"`
	Status    string `json:"status"` // pending, approved, rejected
	CreatedAt string `json:"created_at"`
}
