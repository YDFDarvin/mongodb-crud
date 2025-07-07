package main

import (
	"log"
	"mongodb_crud/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	defer (func() {
		log.Println("Serhio is the best")
		r.Run(":8080")
	})()
	defer (func() {
		r.Use(gin.Logger())
	})()
	defer (func() {
		models.Connect()
	})()

	//models.
}
