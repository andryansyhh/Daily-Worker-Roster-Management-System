package handler

import (
	"net/http"
	"strconv"

	"worker-management/internal/domain/model"
	"worker-management/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ShiftHandler struct {
	Usecase usecase.ShiftUsecase
}

func NewShiftHandler(u usecase.ShiftUsecase) *ShiftHandler {
	return &ShiftHandler{Usecase: u}
}

// POST /admin/shifts
func (h *ShiftHandler) CreateShift(c *gin.Context) {
	var req model.Shift
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	id, err := h.Usecase.CreateShift(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create shift", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Shift created", "shift_id": id})
}

// GET /admin/shifts
func (h *ShiftHandler) GetAllShifts(c *gin.Context) {
	shifts, err := h.Usecase.GetShiftList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get shifts", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"shifts": shifts})
}

// GET /admin/shifts/:id
func (h *ShiftHandler) GetShiftByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	shift, err := h.Usecase.GetShiftByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Shift not found", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"shift": shift})
}

// PUT /admin/shifts/:id
func (h *ShiftHandler) UpdateShift(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req model.Shift
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}
	req.ID = id

	if err := h.Usecase.UpdateShift(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update shift", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shift updated"})
}

// DELETE /admin/shifts/:id
func (h *ShiftHandler) DeleteShift(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.Usecase.DeleteShift(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete shift", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shift deleted"})
}

func (h *ShiftHandler) GetAvailableShifts(c *gin.Context) {
	shifts, err := h.Usecase.GetAvailableShifts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get shifts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"shifts": shifts})
}
