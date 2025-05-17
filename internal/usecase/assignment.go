package usecase

import (
	"errors"
	"worker-management/internal/domain/model"
	"worker-management/internal/repository"
)

type AssignmentUsecase interface {
	GetByDate(date string) ([]model.AssignmentDetail, error)
	GetByWorker(workerID int) ([]model.AssignmentDetail, error)
	ReassignShift(shiftID, newWorkerID int) error
}

type assignmentUsecase struct {
	repo       repository.ShiftRequestRepository
	shiftRepo  repository.ShiftRepository
	workerRepo repository.WorkerRepository
}

func NewAssignmentUsecase(r repository.ShiftRequestRepository, shiftRepo repository.ShiftRepository,
	workerRepo repository.WorkerRepository) AssignmentUsecase {
	return &assignmentUsecase{repo: r, shiftRepo: shiftRepo, workerRepo: workerRepo}

}

func (u *assignmentUsecase) GetByDate(date string) ([]model.AssignmentDetail, error) {
	return u.repo.GetAssignmentsByDate(date)
}

func (u *assignmentUsecase) GetByWorker(workerID int) ([]model.AssignmentDetail, error) {
	return u.repo.GetAssignmentsByWorker(workerID)
}

func (u *assignmentUsecase) ReassignShift(shiftID, newWorkerID int) error {
	// 1. Cek shift 
	shift, err := u.shiftRepo.GetShiftByID(shiftID)
	if err != nil {
		return errors.New("shift not found")
	}

	// 2. Cek worker 
	worker, err := u.workerRepo.GetWorkerByID(newWorkerID)
	if err != nil || worker == nil {
		return errors.New("worker not found")
	}

	// 3. Cek overlap
	overlap, err := u.repo.HasOverlappingShift(newWorkerID, shift.Date, shift.StartTime, shift.EndTime)
	if err != nil {
		return err
	}
	if overlap {
		return errors.New("conflict: new worker has overlapping shift")
	}

	// 4. delete & reassign 
	if err := u.repo.DeleteAssignmentByShiftID(shiftID); err != nil {
		return err
	}
	return u.repo.InsertAssignment(shiftID, newWorkerID)
}
