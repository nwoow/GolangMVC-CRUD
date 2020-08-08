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
 
 This code is using logrus logger to log all the logs in file called gin.log as json.
 
 And whenever it sees any statuscode which is less than:
 
 	statusCode > 399
	
It will capture the logs in databse and you can see that logs using URL/logs 

# Run the code using Docker:

If you wish to tun this code simple run this command to build image:

	docker build -t server .


If you wish to run the code simple run:

	docker run -p 8090:8080 server



