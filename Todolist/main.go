package main

import (
	"todo/Routes"
	"todo/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	repository.Dbconnection()
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/todo", Routes.CreateContent)
	router.PUT("/todo/:id", Routes.CheckboxData)
	router.DELETE("/todo/:id", Routes.DeleteAll)
	router.GET("/todo", Routes.Getalldata)
	router.GET("/activedata", Routes.ActiveData)
	router.GET("/completedata", Routes.CompletedData)
	router.Run(":8000")
}
