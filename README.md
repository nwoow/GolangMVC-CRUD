# Golang MVC-CRUD

To run the Code clone the code:

And then add dependecies this code uses:

	"github.com/Sirupsen/logrus" 
	"github.com/jinzhu/gorm/dialects/sqlite"
  	"github.com/gin-gonic/gin" 
  	"github.com/jinzhu/gorm"
	
You can install all packages before running the code:

After installing packages simply run this:

	go run main.go
 
 
 # Features of this code:
 
 If you run the code you can hit the URL:
 
 	http://localhost:8080/

This will capture a log with error and you can see the logs in this URL:

	http://localhost:8080/logs
	


 # Objective of the code:
 
 In this code you can add a configuration list for logs using:

	http://localhost:8080/logconfig Method:  POST
	BODY: {
	  "title": "Start with Why",
	  "filename": "gin.log",
	  "logtype": "gin.log",
	  "sitename": "gin.log",
	  "loglevel": "info"
	}
 
To list all the confguations file hit this URL:

	http://localhost:8080/logconfiglist Method:  GET

To check all the logs hit this URL:

	http://localhost:8080/logs
	
You can create a configuration file uisng logconfig and you can find a post man coolection in this code postman.json.

Whenver you add a configuration it will take the last filename and read the file and check if level of log is == loglevel then it will push all data to the db and you can see that in logs url.

Feature if you add or update the db of level and filename it will not sync at the moment. 

# Run the code using Docker:

If you wish to tun this code simple run this command to build image:

	docker build -t server .


If you wish to run the code simple run:

	docker run -p 8090:8080 server



