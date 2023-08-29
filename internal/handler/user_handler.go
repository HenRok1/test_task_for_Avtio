package handlers

import (
	"net/http"
	"strconv"

	"github.com/HenRok1/test_task_for_Avito/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	AddUserSegments(c *gin.Context)
	RemoveUserSegments(c *gin.Context)
}

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) AddUserSegments(c *gin.Context) {
	// userID, err := strconv.Atoi(c.Param("user_id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	// 	return
	// }

	var json struct {
		userID       int      `json:"id"`
		segment_name []string `json:"segment_name"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.AddUserSegments(json.userID, json.segment_name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user segments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User segments added seccessfully"})
}

func (h *UserHandler) RemoveUserSegments(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var json struct {
		RemoveSegments []string `json:"removeSegments"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.RemoveUserSegments(userID, json.RemoveSegments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user segments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User segments updated successfully"})

}
