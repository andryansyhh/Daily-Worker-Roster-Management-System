package model

type ShiftFilter struct {
	Date  string `form:"date"`
	Role  string `form:"role"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}
