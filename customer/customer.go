package customer

import (
	"../database"
	"../models"
	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	database.CONNECTION.Find(&customers)
	var customersResponse = models.CategorizeCustomersByPhoneNumbers(customers)
	c.JSON(200, &customersResponse)
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	database.CONNECTION.Find(&customer, id)
	c.JSON(200, customer)
}

func AddCustomer(c *gin.Context) {
	var customer models.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.Status(400)
		return
	}
	database.CONNECTION.Create(&customer)
}
