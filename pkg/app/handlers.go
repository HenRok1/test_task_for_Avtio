package pkg

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID int `json:"user_id"`
	// Name string `json:"name"`
	// Другие поля пользователя
}

type Segment struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Users []int  `json:"users_id"`
	// Другие поля сегмента
}

var users map[int]*User
var segments map[int]*Segment

func main() {
	users = make(map[int]*User)
	segments = make(map[int]*Segment)

	router := gin.Default()

	// Создание пользователя
	router.POST("/users", createUser)

	// Получение информации о пользователе
	router.GET("/users/:id", getUser)

	// Изменение пользователя
	router.PUT("/users/:id", updateUser)

	// Удаление пользователя
	router.DELETE("/users/:id", deleteUser)

	// Создание сегмента
	router.POST("/segments", createSegment)

	// Получение информации о сегменте
	router.GET("/segments/:id", getSegment)

	// Изменение сегмента
	router.PUT("/segments/:id", updateSegment)

	// Удаление сегмента
	router.DELETE("/segments/:id", deleteSegment)

	// Добавление пользователя в сегмент
	router.POST("/segments/:id/users/:userID", addSegmentUser)

	// Удаление пользователя из сегмента
	router.DELETE("/segments/:id/users/:userID", removeSegmentUser)

	router.Run(":8080")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users[user.ID] = &user
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func getUser(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, ok := users[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Обработчики для обновления и удаления пользователя аналогичны

func createSegment(c *gin.Context) {
	var segment Segment
	if err := c.ShouldBindJSON(&segment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	segments[segment.ID] = &segment
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func getSegment(c *gin.Context) {
	segmentID := c.Param("id")
	id, err := strconv.Atoi(segmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	segment, ok := segments[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Segment not found"})
		return
	}

	c.JSON(http.StatusOK, segment)
}

// Обработчики для обновления и удаления сегмента аналогичны

func addSegmentUser(c *gin.Context) {
	segmentID := c.Param("id")
	id, err := strconv.Atoi(segmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	segment, ok := segments[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Segment not found"})
		return
	}

	userID := c.Param("userID")
	new_userID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	segment.Users = append(segment.Users, new_userID)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
