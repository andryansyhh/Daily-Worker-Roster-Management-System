package repository

import (
	"database/sql"
	"worker-management/internal/domain/model"
)

type ShiftRepository interface {
	CreateShift(shift model.Shift) (int, error)
	GetShiftByID(id int) (*model.Shift, error)
	GetAllShifts() ([]model.Shift, error)
	UpdateShift(shift model.Shift) error
	DeleteShift(id int) error
	GetAvailableShifts() ([]model.Shift, error)
}

type shiftRepository struct {
	db *sql.DB
}

func NewShiftRepository(db *sql.DB) ShiftRepository {
	return &shiftRepository{db: db}
}

func (r *shiftRepository) CreateShift(s model.Shift) (int, error) {
	query := `INSERT INTO shifts (date, start_time, end_time, role, location) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, s.Date, s.StartTime, s.EndTime, s.Role, s.Location)
	if err != nil {
		return 0, err
	}
	insertedID, _ := result.LastInsertId()
	return int(insertedID), nil
}

func (r *shiftRepository) GetShiftByID(id int) (*model.Shift, error) {
	var s model.Shift
	query := `SELECT id, date, start_time, end_time, role, location FROM shifts WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&s.ID, &s.Date, &s.StartTime, &s.EndTime, &s.Role, &s.Location)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *shiftRepository) GetAllShifts() ([]model.Shift, error) {
	query := `SELECT id, date, start_time, end_time, role, location FROM shifts ORDER BY date, start_time`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []model.Shift
	for rows.Next() {
		var s model.Shift
		if err := rows.Scan(&s.ID, &s.Date, &s.StartTime, &s.EndTime, &s.Role, &s.Location); err != nil {
			return nil, err
		}
		shifts = append(shifts, s)
	}
	return shifts, nil
}

func (r *shiftRepository) UpdateShift(s model.Shift) error {
	query := `UPDATE shifts SET date = ?, start_time = ?, end_time = ?, role = ?, location = ? WHERE id = ?`
	_, err := r.db.Exec(query, s.Date, s.StartTime, s.EndTime, s.Role, s.Location, s.ID)
	return err
}

func (r *shiftRepository) DeleteShift(id int) error {
	_, err := r.db.Exec(`DELETE FROM shifts WHERE id = ?`, id)
	return err
}

func (r *shiftRepository) GetAvailableShifts() ([]model.Shift, error) {
	query := `
	SELECT s.id, s.date, s.start_time, s.end_time, s.role, s.location
	FROM shifts s
	LEFT JOIN assignments a ON s.id = a.shift_id
	WHERE a.shift_id IS NULL
	ORDER BY s.date, s.start_time
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []model.Shift
	for rows.Next() {
		var s model.Shift
		if err := rows.Scan(&s.ID, &s.Date, &s.StartTime, &s.EndTime, &s.Role, &s.Location); err != nil {
			return nil, err
		}
		shifts = append(shifts, s)
	}
	return shifts, nil
}
