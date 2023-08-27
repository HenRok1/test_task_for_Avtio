package handler

import (
	"net/http"

	"github.com/HenRok1/test_task_for_Avito/internal/db"
	"github.com/HenRok1/test_task_for_Avito/internal/entity"
	"github.com/gin-gonic/gin"
)

func CreateSegment(c *gin.Context) {
	var segment entity.Segment
	if err := c.ShouldBindJSON(&segment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert segment into the database
	_, err := db.GetDB().Exec("INSERT INTO segments (name) VALUES ($1)", segment.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create segment"})
		return
	}

	c.JSON(http.StatusCreated, segment)
}

func DeleteSegment(c *gin.Context) {
	name := c.Param("name")

	// Delete segment from the database
	_, err := db.GetDB().Exec("DELETE FROM segments WHERE name = $1", name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete segment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segment deleted"})
}
