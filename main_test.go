package main

import (
	"./models"
	"./utils"
	"testing"
)

func TestCategorizeCustomersByPhoneNumbers(t *testing.T) {

	// test that this person is from Morocco
	var person models.Customer
	person.Name = "Test Person"
	person.Phone = "(212) 691933626"

	var listOfPersons []models.Customer
	listOfPersons = append(listOfPersons, person)

	var customerResponses = utils.CategorizeCustomersByPhoneNumbers(listOfPersons)

	if customerResponses[0].Country.Name != "Morocco" {
		t.Error("Expected Phone to be Morocco")
	}
}
