package Router

import (
	"wordplay/Logger"

	controllers "wordplay/Controllers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Router(logger logrus.FieldLogger) {
	r := gin.New()
	r.Use(Logger.Logger(logger), gin.Recovery())
	// r.POST("/books", controllers.CreateBook)
	r.GET("/logs", controllers.Listerror) // new
	// r.GET("/books/:id", controllers.FindBook)     // new
	// r.PATCH("/books/:id", controllers.UpdateBook) // new
	r.Run(":8080")
}
