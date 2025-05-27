package ginrestapi

import "github.com/gin-gonic/gin"

func Init() {
	// Create a new Gin router
	g := gin.Default()
	g.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	g.Run(":8080") // Start the server on port

}
