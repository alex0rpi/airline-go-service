package main

import (
	"go-fleet-management/internal/api/router"
	"go-fleet-management/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database connection
    db.InitDB()

    // Set up the Gin router
    r := gin.Default()

    // Set up the routes
    router.SetupRoutes(r)

    // Start the HTTP server
    r.Run(":8080") // Listen and serve on port 8080
}