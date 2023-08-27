package handler

import (
	"net/http"

	"github.com/HenRok1/test_task_for_Avito/internal/db"
	"github.com/HenRok1/test_task_for_Avito/internal/entity"

	"github.com/gin-gonic/gin"
)

func AddUserSegments(c *gin.Context) {
	userID := c.Param("userID")

	var userSegments entity.UserSegment
	if err := c.ShouldBindJSON(&userSegments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add and remove user segments in the database
	tx, err := db.GetDB().Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	// Delete existing user segments
	_, err = tx.Exec("DELETE FROM user_segments WHERE user_id = $1", userID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user segments"})
		return
	}

	// Insert new user segments
	for _, segmentName := range userSegments.SegmentName {
		_, err = tx.Exec("INSERT INTO user_segments (user_id, segment_name) VALUES ($1, $2)", userID, segmentName)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user segments"})
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "User segments updated"})
}

func GetUserSegments(c *gin.Context) {
	userID := c.Param("userID")

	// Retrieve active segments for the user from the database
	rows, err := db.GetDB().Query("SELECT s.id, s.slug FROM segments s INNER JOIN user_segments us ON s.id = us.segment_id WHERE us.user_id = $1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user segments"})
		return
	}
	defer rows.Close()

	segments := []entity.Segment{}
	for rows.Next() {
		var segment entity.Segment
		err := rows.Scan(&segment.ID, &segment.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user segments"})
			return
		}
		segments = append(segments, segment)
	}

	c.JSON(http.StatusOK, segments)
}
