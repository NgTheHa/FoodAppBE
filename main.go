package main

import (
	"github.com/gin-gonic/gin"
	"go/foodappbe/Infrastructure"
	"log"
)

func main() {

	router := gin.Default()

	Infrastructure.ConnectDatabase()
	log.Println("Server is running...")

	//server port
	router.Run(":8080")
}
