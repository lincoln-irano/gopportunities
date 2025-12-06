package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router with *GIN*
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run server on port http://localhost:8080
	router.Run(":8080")
}
