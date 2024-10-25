package routes

import (
	"blaster_balu/controllers"
	"blaster_balu/repository"
	"blaster_balu/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, db *sql.DB) {
	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Set up routes
	SignUpRoutes(r, userController)
}
