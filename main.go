package main

import (
	"./customer"
	"./database"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initDatabase() {
	var err error
	database.CONNECTION, err = gorm.Open(sqlite.Open("sample.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected successfully")
}

func setupRoutes() {
	r := gin.Default()
	r.GET("/customers", customer.GetCustomers)
	r.GET("/customers/:id", customer.GetCustomer)
	r.POST("/customers/add", customer.AddCustomer)

	err := r.Run()
	if err != nil {
		panic("could not run gin app")
		return
	}
}

func main() {
	initDatabase()
	setupRoutes()
}
