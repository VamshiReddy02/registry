package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vamshireddy02/registry/middleware"
	"github.com/vamshireddy02/registry/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize the Gin router
	router := gin.New()

	// Register the logger middleware
	router.Use(gin.Logger())

	// Register the metrics middleware
	router.Use(middleware.MetricsMiddleware())

	// Set up Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Set up the routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// Run the server
	router.Run(":" + port)
}
