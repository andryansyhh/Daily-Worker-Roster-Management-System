package handler

import (
	"net/http"
	"worker-management/internal/domain/model"
	"worker-management/internal/usecase"

	"github.com/gin-gonic/gin"
)

type WorkerHandler struct {
	usecase usecase.WorkerUsecase
}

func NewWorkerHandler(u usecase.WorkerUsecase) *WorkerHandler {
	return &WorkerHandler{usecase: u}
}

func (h *WorkerHandler) Register(c *gin.Context) {
	var req model.Worker
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	id, err := h.usecase.Register(&req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Worker registered successfully",
		"worker_id": id,
	})
}
