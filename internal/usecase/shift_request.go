package usecase

import (
	"errors"
	"worker-management/internal/domain/model"
	"worker-management/internal/repository"
)

type ShiftRequestUsecase interface {
	RequestShift(workerID, shiftID int) error
	GetMyRequests(workerID int) ([]model.ShiftRequest, error)
	ApproveRequest(requestID int) error
	RejectRequest(requestID int) error
}

type shiftRequestUsecase struct {
	shiftRequestRepo repository.ShiftRequestRepository
	shiftRepo        repository.ShiftRepository
}

func NewShiftRequestUsecase(shiftRequestRepo repository.ShiftRequestRepository, shiftRepo repository.ShiftRepository) ShiftRequestUsecase {
	return &shiftRequestUsecase{
		shiftRequestRepo: shiftRequestRepo,
		shiftRepo:        shiftRepo,
	}
}

func (u *shiftRequestUsecase) RequestShift(workerID, shiftID int) error {

	// 1. get if shift already assign
	assigned, err := u.shiftRequestRepo.IsShiftAlreadyAssigned(shiftID)
	if err != nil {

		return err
	}

	if assigned {
		return errors.New("shift already assigned to someone")
	}

	// 2. get detail shift
	shift, err := u.shiftRepo.GetShiftByID(shiftID)
	if err != nil {

		return errors.New("shift not found")
	}

	// 3. Cek overlap
	overlap, err := u.shiftRequestRepo.HasOverlappingShift(workerID, shift.Date, shift.StartTime, shift.EndTime)
	if err != nil {

		return err
	}
	if overlap {
		return errors.New("shift overlaps with existing assignment")
	}

	// 4. Cek shift same date
	countToday, err := u.shiftRequestRepo.CountShiftsThisWeek(workerID, shift.Date)
	if err != nil {

		return err
	}
	if countToday >= 5 {
		return errors.New("max 5 shifts allowed per week")
	}

	// 5. save shift request
	return u.shiftRequestRepo.CreateRequest(shiftID, workerID)
}

func (u *shiftRequestUsecase) GetMyRequests(workerID int) ([]model.ShiftRequest, error) {
	return u.shiftRequestRepo.GetRequestsByWorker(workerID)
}

func (u *shiftRequestUsecase) ApproveRequest(requestID int) error {
	// 1. get data request
	req, err := u.shiftRequestRepo.GetRequestByID(requestID)
	if err != nil {
		return err
	}
	if req.Status != "pending" {
		return errors.New("request is not in pending status")
	}

	// 2. Insert assignments
	err = u.shiftRequestRepo.InsertAssignment(req.ShiftID, req.WorkerID)
	if err != nil {
		return err
	}

	// 3. Update status request
	return u.shiftRequestRepo.UpdateStatusRequest(requestID, "approved")
}

func (u *shiftRequestUsecase) RejectRequest(requestID int) error {
	req, err := u.shiftRequestRepo.GetRequestByID(requestID)
	if err != nil {
		return err
	}
	if req.Status != "pending" {
		return errors.New("request is not in pending status")
	}
	return u.shiftRequestRepo.UpdateStatusRequest(requestID, "rejected")
}
