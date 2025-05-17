package repository

import (
	"database/sql"
	"fmt"
	"worker-management/internal/domain/model"
)

type ShiftRequestRepository interface {
	CreateRequest(shiftID, workerID int) error
	GetRequestsByWorker(workerID int) ([]model.ShiftRequest, error)
	IsShiftAlreadyAssigned(shiftID int) (bool, error)
	HasOverlappingShift(workerID int, date, start, end string) (bool, error)
	CountShiftsThisWeek(workerID int, weekStart string) (int, error)
	UpdateStatusRequest(requestID int, status string) error
	GetRequestByID(id int) (*model.ShiftRequest, error)
	InsertAssignment(shiftID, workerID int) error
	GetAssignmentsByDate(date string) ([]model.AssignmentDetail, error)
	GetAssignmentsByWorker(workerID int) ([]model.AssignmentDetail, error)
	DeleteAssignmentByShiftID(shiftID int) error
}

type shiftRequestRepository struct {
	db *sql.DB
}

func NewShiftRequestRepository(db *sql.DB) ShiftRequestRepository {
	return &shiftRequestRepository{db: db}
}

func (r *shiftRequestRepository) CreateRequest(shiftID, workerID int) error {
	query := `
		INSERT INTO shift_requests (shift_id, worker_id, status, created_at)
		VALUES (?, ?, 'pending', datetime('now'))
	`
	_, err := r.db.Exec(query, shiftID, workerID)
	if err != nil {
		fmt.Println("Insert error:", err)
		return err
	}

	return nil
}

func (r *shiftRequestRepository) GetRequestsByWorker(workerID int) ([]model.ShiftRequest, error) {
	query := `
		SELECT id, shift_id, worker_id, status, created_at
		FROM shift_requests
		WHERE worker_id = ?
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query, workerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []model.ShiftRequest
	for rows.Next() {
		var req model.ShiftRequest
		if err := rows.Scan(&req.ID, &req.ShiftID, &req.WorkerID, &req.Status, &req.CreatedAt); err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}
	return requests, nil
}

func (r *shiftRequestRepository) IsShiftAlreadyAssigned(shiftID int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM assignments WHERE shift_id = ?`
	err := r.db.QueryRow(query, shiftID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *shiftRequestRepository) HasOverlappingShift(workerID int, date, start, end string) (bool, error) {
	query := `
	SELECT COUNT(*)
	FROM assignments a
	JOIN shifts s ON a.shift_id = s.id
	WHERE a.worker_id = ?
	  AND s.date = ?
	  AND (
	    (s.start_time < ? AND s.end_time > ?) OR
	    (s.start_time >= ? AND s.start_time < ?)
	  )
	`
	var count int
	err := r.db.QueryRow(query, workerID, date, end, start, start, end).Scan(&count)
	return count > 0, err
}

func (r *shiftRequestRepository) CountShiftsThisWeek(workerID int, weekStart string) (int, error) {
	query := `
	SELECT COUNT(*)
	FROM assignments a
	JOIN shifts s ON a.shift_id = s.id
	WHERE a.worker_id = ?
	  AND s.date BETWEEN ? AND DATE(?, '+6 days')
	`
	var count int
	err := r.db.QueryRow(query, workerID, weekStart, weekStart).Scan(&count)
	return count, err
}

func (r *shiftRequestRepository) UpdateStatusRequest(requestID int, status string) error {
	query := `UPDATE shift_requests SET status = ? WHERE id = ?`
	_, err := r.db.Exec(query, status, requestID)
	return err
}

func (r *shiftRequestRepository) GetRequestByID(id int) (*model.ShiftRequest, error) {
	query := `SELECT id, shift_id, worker_id, status, created_at FROM shift_requests WHERE id = ?`
	var req model.ShiftRequest
	err := r.db.QueryRow(query, id).Scan(
		&req.ID, &req.ShiftID, &req.WorkerID, &req.Status, &req.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *shiftRequestRepository) InsertAssignment(shiftID, workerID int) error {
	query := `INSERT INTO assignments (shift_id, worker_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, shiftID, workerID)
	return err
}

func (r *shiftRequestRepository) GetAssignmentsByDate(date string) ([]model.AssignmentDetail, error) {
	query := `
	SELECT a.shift_id, a.worker_id, s.date, s.start_time, s.end_time, s.role, s.location
	FROM assignments a
	JOIN shifts s ON a.shift_id = s.id
	WHERE s.date = ?
	ORDER BY s.start_time
	`
	rows, err := r.db.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.AssignmentDetail
	for rows.Next() {
		var a model.AssignmentDetail
		if err := rows.Scan(&a.ShiftID, &a.WorkerID, &a.Date, &a.StartTime, &a.EndTime, &a.Role, &a.Location); err != nil {
			return nil, err
		}
		results = append(results, a)
	}
	return results, nil
}

func (r *shiftRequestRepository) GetAssignmentsByWorker(workerID int) ([]model.AssignmentDetail, error) {
	query := `
	SELECT a.shift_id, a.worker_id, s.date, s.start_time, s.end_time, s.role, s.location
	FROM assignments a
	JOIN shifts s ON a.shift_id = s.id
	WHERE a.worker_id = ?
	ORDER BY s.date, s.start_time
	`
	rows, err := r.db.Query(query, workerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.AssignmentDetail
	for rows.Next() {
		var a model.AssignmentDetail
		if err := rows.Scan(&a.ShiftID, &a.WorkerID, &a.Date, &a.StartTime, &a.EndTime, &a.Role, &a.Location); err != nil {
			return nil, err
		}
		results = append(results, a)
	}
	return results, nil
}

func (r *shiftRequestRepository) DeleteAssignmentByShiftID(shiftID int) error {
	_, err := r.db.Exec(`DELETE FROM assignments WHERE shift_id = ?`, shiftID)
	return err
}
