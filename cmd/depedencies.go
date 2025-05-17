package cmd

import (
	"log"
	"worker-management/internal/handler"
	"worker-management/internal/repository"
	"worker-management/internal/usecase"
)

type Dependency struct {
	Config               *Config
	ShiftHandler         *handler.ShiftHandler
	ShiftReqHandler      *handler.RequestHandler
	AssignmentReqHandler *handler.AssignmentHandler
	AuthHandler          *handler.AuthHandler
	WorkerHandler        *handler.WorkerHandler
}

func InitDependencies() *Dependency {
	cfg, err := Load()
	if err != nil {
		log.Println("error to load")
	}

	// External dependencies
	dbConn := NewClientDatabase()

	// shift module
	shiftRepository := repository.NewShiftRepository(dbConn)
	shiftUsecase := usecase.NewShiftUsecase(shiftRepository)
	shiftHandler := handler.NewShiftHandler(shiftUsecase)

	// jobs module
	shiftRequetsRepository := repository.NewShiftRequestRepository(dbConn)
	shiftRequetsUsecase := usecase.NewShiftRequestUsecase(shiftRequetsRepository, shiftRepository)
	shiftRequetsHandler := handler.NewRequestHandler(shiftRequetsUsecase)

	// worker module
	workerRepository := repository.NewWorkerRequestRepository(dbConn)
	workerUsecase := usecase.NewWorkerUsecase(workerRepository)
	workerHandler := handler.NewWorkerHandler(workerUsecase)

	// assignment module
	assignmentUsecase := usecase.NewAssignmentUsecase(shiftRequetsRepository, shiftRepository, workerRepository)
	assignmentHandler := handler.NewAssignmentHandler(assignmentUsecase)

	authHandler := handler.NewAuthHandler()

	return &Dependency{
		Config:               cfg,
		ShiftHandler:         shiftHandler,
		ShiftReqHandler:      shiftRequetsHandler,
		AssignmentReqHandler: assignmentHandler,
		AuthHandler:          authHandler,
		WorkerHandler:        workerHandler,
	}
}
