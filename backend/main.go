package main

import (
	"example.com/sa-68-example/config"
	"example.com/sa-68-example/controller" // import controller package
	"example.com/sa-68-example/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

const PORT = "8000"

func main() {

	// open connection database
	config.ConnectionDB()

	// Generate databases
	config.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")

	{
		router.Use(middlewares.Authorizes())

		// LibrarianMember Routes
		router.PUT("/librarianmember/:memberID", librarianmembers.Update)    // ใช้ controller.Update
		router.GET("/librarianmembers", librarianmembers.GetAll)             // ใช้ controller.GetAll
		router.GET("/librarianmember/:memberID", librarianmembers.Get)       // ใช้ controller.Get
		router.DELETE("/librarianmember/:memberID", librarianmembers.Delete) // ใช้ controller.Delete
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	// Run the server
	r.Run("localhost:" + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
