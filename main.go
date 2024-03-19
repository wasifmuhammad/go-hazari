// steps
// commands
// go mod init golang-crud-gin
// go get github.com/gin-gonic/gin
// go get gorm.io/gorm
// go get gorm.io/driver/postgres
// go get github.com/mattn/go-isatty@v0.0.19

// to run

// docker-compose up --build
// go run main.go
package main

import (
	"golang-crud-gin/config"
	"golang-crud-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router:=gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8000")
}



//  go run main.go   
