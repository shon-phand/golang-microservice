package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/bookstore_users-api/controllers/configuration"
)

var (
	r = gin.Default()
)

func StartApp() {
	fmt.Println("webserver started")
	config, _ := configuration.LoadConfiguration("/home/shon/Documents/Microservice/golang-microservice/src/github.com/shon-phand/bookstore_users-api/app/webserver_properties.json")
	mapUrls()

	r.Run(config.Host + ":" + config.Port)
}
