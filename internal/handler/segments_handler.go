package handlers

import (
	"net/http"

	"github.com/HenRok1/test_task_for_Avito/internal/services"
	"github.com/gin-gonic/gin"
)

type SegmentHandler struct {
	segmentService *services.SegmentService
}

func NewSegmentHandler(segmentService *services.SegmentService) *SegmentHandler {
	return &SegmentHandler{
		segmentService: segmentService,
	}
}

func (h *SegmentHandler) CreateSegment(c *gin.Context) {
	var json struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.segmentService.CreateSegment(json.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create segment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Segment created successfully"})
}

func (h *SegmentHandler) DeleteSegment(c *gin.Context) {
	name := c.Param("name")

	if err := h.segmentService.DeleteSegment(name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete segment"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segment deleted successfully"})
}
