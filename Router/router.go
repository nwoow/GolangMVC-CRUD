package Router

import (
	controllers "wordplay/Controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.POST("/logconfig", controllers.Logconfigure)
	r.GET("/logs", controllers.Listerror)              // new
	r.GET("/logconfiglist", controllers.Logconfiglist) // new
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// r.GET("/books/:id", controllers.FindBook)     // new
	// r.PATCH("/books/:id", controllers.UpdateBook) // new
	r.Run(":8080")
}
