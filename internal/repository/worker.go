package repository

import (
	"database/sql"
	"worker-management/internal/domain/model"
)

type WorkerRepository interface {
	GetWorkerByID(id int) (*model.Worker, error)
	Create(worker *model.Worker) (int64, error)
}

type workerRepository struct {
	db *sql.DB
}

func NewWorkerRequestRepository(db *sql.DB) WorkerRepository {
	return &workerRepository{db: db}
}

func (r *workerRepository) GetWorkerByID(id int) (*model.Worker, error) {
	query := `SELECT id, name, email FROM workers WHERE id = ?`
	var w model.Worker
	err := r.db.QueryRow(query, id).Scan(&w.ID, &w.Name, &w.Email)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *workerRepository) Create(worker *model.Worker) (int64, error) {
	query := `INSERT INTO workers (name, email) VALUES (?, ?)`
	res, err := r.db.Exec(query, worker.Name, worker.Email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
