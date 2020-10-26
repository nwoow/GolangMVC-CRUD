package main

import (
	models "wordplay/Models"
	"wordplay/Router"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	models.ConnectDataBase() // new
	Router.Router()

}
