package router

import (
	"worker-management/internal/handler"
	"worker-management/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, shiftHandler *handler.ShiftHandler,
	requestHandler *handler.RequestHandler, assignmentHandler *handler.AssignmentHandler, authHandler handler.AuthHandler, workerHandler *handler.WorkerHandler) {

	// register
	r.POST("/worker/register", workerHandler.Register)
	//login
	r.POST("/auth/login", authHandler.Login)

	admin := r.Group("/admin", middleware.JWTAuthMiddleware(), middleware.RequireRole("admin"))
	{
		admin.POST("/shifts", shiftHandler.CreateShift)
		admin.GET("/shifts", shiftHandler.GetAllShifts)
		admin.GET("/shifts/:id", shiftHandler.GetShiftByID)
		admin.PUT("/shifts/:id", shiftHandler.UpdateShift)
		admin.DELETE("/shifts/:id", shiftHandler.DeleteShift)

		admin.PUT("/requests/:id/approve", requestHandler.ApproveRequest)
		admin.PUT("/requests/:id/reject", requestHandler.RejectRequest)

		admin.GET("/assignments", assignmentHandler.GetByDate)
		admin.GET("/assignments/worker/:id", assignmentHandler.GetByWorker)
		admin.PUT("/assignments/:shift_id/reassign", assignmentHandler.ReassignShift)
	}

	worker := r.Group("/worker", middleware.JWTAuthMiddleware(), middleware.RequireRole("worker"))
	{
		worker.GET("/shifts/available", shiftHandler.GetAvailableShifts)
		worker.GET("/requests", requestHandler.GetMyRequests)
		worker.POST("/shifts/:id/request", requestHandler.RequestShift)
		worker.GET("/assignments", assignmentHandler.GetMine)
	}

}
