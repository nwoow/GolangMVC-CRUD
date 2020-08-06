package Router

import (
	"wordplay/Logger"

	controllers "wordplay/Controllers"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Router(logger logrus.FieldLogger) {
	r := gin.New()
	r.Use(Logger.Logger(logger), gin.Recovery())
	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.FindBooks)        // new
	r.GET("/books/:id", controllers.FindBook)     // new
	r.PATCH("/books/:id", controllers.UpdateBook) // new
	r.Run(":8080")
}
