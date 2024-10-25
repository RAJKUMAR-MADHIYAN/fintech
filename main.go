package main

import (
	"blaster_balu/database"
	"blaster_balu/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.GetDB()
	defer db.Close()
	routes.Setup(router, db)

	// Start server
	router.Run(":8080")
}
