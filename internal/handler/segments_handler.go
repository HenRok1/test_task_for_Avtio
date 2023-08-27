package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/HenRok1/test_task_for_Avito/internal/services"
)

type SegmentHandler struct {
	SegmentService services.SegmentService
}

func (h *SegmentHandler) CreateSegment(c *gin.Context) {
	var input struct {
		Name string `json:"Name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	segment, err := h.SegmentService.CreateSegment(input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create segment"})
		return
	}

	c.JSON(http.StatusCreated, segment)
}
