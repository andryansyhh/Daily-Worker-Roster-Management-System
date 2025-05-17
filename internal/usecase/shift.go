package usecase

import (
	"worker-management/internal/domain/model"
	"worker-management/internal/repository"
)

type ShiftUsecase interface {
	CreateShift(input model.Shift) (int, error)
	GetShiftList() ([]model.Shift, error)
	GetShiftByID(id int) (*model.Shift, error)
	UpdateShift(input model.Shift) error
	DeleteShift(id int) error
	GetAvailableShifts() ([]model.Shift, error)
}

type shiftUsecase struct {
	shiftRepo repository.ShiftRepository
}

func NewShiftUsecase(r repository.ShiftRepository) ShiftUsecase {
	return &shiftUsecase{
		shiftRepo: r,
	}
}

func (u *shiftUsecase) CreateShift(s model.Shift) (int, error) {
	// Tambah validasi basic kalau perlu (misal waktu kosong)

	// Simpan shift
	return u.shiftRepo.CreateShift(s)
}

func (u *shiftUsecase) GetShiftList() ([]model.Shift, error) {
	return u.shiftRepo.GetAllShifts()
}

func (u *shiftUsecase) GetShiftByID(id int) (*model.Shift, error) {
	return u.shiftRepo.GetShiftByID(id)
}

func (u *shiftUsecase) UpdateShift(s model.Shift) error {
	return u.shiftRepo.UpdateShift(s)
}

func (u *shiftUsecase) DeleteShift(id int) error {
	return u.shiftRepo.DeleteShift(id)
}

func (u *shiftUsecase) GetAvailableShifts() ([]model.Shift, error) {
	return u.shiftRepo.GetAvailableShifts()
}
