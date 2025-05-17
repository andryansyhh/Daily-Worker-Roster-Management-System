package usecase

import (
	"worker-management/internal/domain/model"
	"worker-management/internal/repository"
)

type WorkerUsecase interface {
	Register(worker *model.Worker) (int64, error)
}

type workerUsecase struct {
	repo repository.WorkerRepository
}

func NewWorkerUsecase(repo repository.WorkerRepository) WorkerUsecase {
	return &workerUsecase{repo: repo}
}

func (u *workerUsecase) Register(worker *model.Worker) (int64, error) {
	return u.repo.Create(worker)
}
