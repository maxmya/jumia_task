package main

import (
	"./customer"
	"./database"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware : allow CORS access from all
func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Next()
	}
}

func setupRoutes() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/customers", customer.GetCustomers)
	router.GET("/customers/:id", customer.GetCustomer)
	router.POST("/customers/add", customer.AddCustomer)

	err := router.Run()
	if err != nil {
		panic("could not run gin app")
		return
	}
}

func main() {
	database.InitDatabase()
	setupRoutes()
}
