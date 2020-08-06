package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	rxURL = regexp.MustCompile(`^/regexp\d*`)
)

func main() {
	// models.ConnectDataBase() // new
	// logger := logrus.New()
	// logger.Formatter = &logrus.JSONFormatter{}

	// logger.SetOutput(os.Stdout)
	// file, err := os.OpenFile("gin.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	// if err != nil {
	// 	logger.Fatal(err)
	// }

	// defer file.Close()
	// logger.SetOutput(file)

	// Router.Router(logger)

	// Open our jsonFile
	jsonFile, err := os.Open("gin.log")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["users"])

}
