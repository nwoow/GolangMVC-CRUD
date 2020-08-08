package main

import (
	"os"
	"time"
	models "wordplay/Models"
	"wordplay/Router"
	"wordplay/Tali"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

func log_event() {
	c := time.Tick(5 * time.Second)
	var Logconfigure []models.Logconfigure
	models.DB.Last(&Logconfigure)
	var file string
	if err := models.DB.Last(&Logconfigure).Error; err != nil {
		file = "gin.log"
	} else {
		file = Logconfigure[0].Filename
	}
	println(file)
	for _ = range c {
		Tali.Tali1(file)
	}

}
func main() {

	models.ConnectDataBase() // new
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	logger.SetOutput(os.Stdout)
	file, err := os.OpenFile("gin.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}

	defer file.Close()
	logger.SetOutput(file)
	go log_event()
	Router.Router(logger)

}
