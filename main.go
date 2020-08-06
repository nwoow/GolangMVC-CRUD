package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
	models "wordplay/Models"
	"wordplay/Router"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

var (
	rxURL = regexp.MustCompile(`^/regexp\d*`)
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
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
	Router.Router(logger)

}
