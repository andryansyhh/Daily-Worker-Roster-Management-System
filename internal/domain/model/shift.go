package model

type Shift struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`        // Format: YYYY-MM-DD
	StartTime string `json:"start_time"`  // Format: HH:MM (in UTC)
	EndTime   string `json:"end_time"`    // Format: HH:MM (in UTC)
	Role      string `json:"role"`
	Location  string `json:"location,omitempty"`
}
