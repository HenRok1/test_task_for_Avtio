package handlers

import "github.com/HenRok1/test_task_for_Avito/internal/services"

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
