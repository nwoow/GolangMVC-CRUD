package main

import (
	"os"
	"regexp"
	models "wordplay/Models"
	"wordplay/Router"

	"github.com/Sirupsen/logrus"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	rxURL = regexp.MustCompile(`^/regexp\d*`)
)

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

	Router.Router(logger)

}
