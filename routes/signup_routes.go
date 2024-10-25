package routes

import (
	"blaster_balu/controllers"

	"github.com/gin-gonic/gin"
)

func SignUpRoutes(r *gin.Engine, userController *controllers.UserController) {
	api := r.Group("/api")
	{
		api.POST("/signup", userController.Signup)
	}
}
