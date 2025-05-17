package handler

import (
	"net/http"
	"strconv"
	"worker-management/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Usecase usecase.ShiftRequestUsecase
}

func NewRequestHandler(u usecase.ShiftRequestUsecase) *RequestHandler {
	return &RequestHandler{Usecase: u}
}

func (h *RequestHandler) RequestShift(c *gin.Context) {
	shiftID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid shift ID"})
		return
	}

	workerID, _ := c.Get("user_id")

	err = h.Usecase.RequestShift(workerID.(int), shiftID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to request shift",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shift request created"})
}

func (h *RequestHandler) GetMyRequests(c *gin.Context) {
	workerID, _ := c.Get("user_id")

	requests, err := h.Usecase.GetMyRequests(workerID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get requests",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

func (h *RequestHandler) ApproveRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.Usecase.ApproveRequest(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Approval failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Request approved"})
}

func (h *RequestHandler) RejectRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.Usecase.RejectRequest(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Rejection failed", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Request rejected"})
}
