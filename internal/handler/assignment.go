package handler

import (
	"log"
	"net/http"
	"strconv"
	"worker-management/internal/domain/model"
	"worker-management/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AssignmentHandler struct {
	Usecase usecase.AssignmentUsecase
}

func NewAssignmentHandler(u usecase.AssignmentUsecase) *AssignmentHandler {
	return &AssignmentHandler{Usecase: u}
}

func (h *AssignmentHandler) GetByDate(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing date query"})
		return
	}
	assignments, err := h.Usecase.GetByDate(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"assignments": assignments})
}

func (h *AssignmentHandler) GetByWorker(c *gin.Context) {
	workerID, _ := strconv.Atoi(c.Param("id"))
	assignments, err := h.Usecase.GetByWorker(workerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"assignments": assignments})
}

func (h *AssignmentHandler) GetMine(c *gin.Context) {
	workerID := 1 // dummy
	assignments, err := h.Usecase.GetByWorker(workerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"assignments": assignments})
}

func (h *AssignmentHandler) ReassignShift(c *gin.Context) {
	shiftID, err := strconv.Atoi(c.Param("shift_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shift_id"})
		return
	}

	var req model.Worker
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	if err := h.Usecase.ReassignShift(shiftID, req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[REASSIGN] Shift %d assigned to worker %d\n", shiftID, req.ID)
	c.JSON(http.StatusOK, gin.H{"message": "Shift reassigned"})
}
