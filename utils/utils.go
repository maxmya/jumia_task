package utils

import (
	"../models"
	"fmt"
	"regexp"
)

func CategorizeCustomersByPhoneNumbers(customers []models.Customer) []models.CustomersResponse {

	var customersResponse []models.CustomersResponse

	var cameronCustomers models.CustomersResponse
	var cameronCountry models.Country
	cameronCountry.Name = "Cameron"
	cameronCountry.Code = "+237"
	cameronCustomers.Country = cameronCountry

	var ethiopiaCustomers models.CustomersResponse
	var ethiopiaCountry models.Country
	ethiopiaCountry.Name = "Ethiopia"
	ethiopiaCountry.Code = "+251"
	ethiopiaCustomers.Country = ethiopiaCountry

	var moroccoCustomers models.CustomersResponse
	var moroccoCountry models.Country
	moroccoCountry.Name = "Morocco"
	moroccoCountry.Code = "+212"
	moroccoCustomers.Country = moroccoCountry

	var mozambiqueCustomers models.CustomersResponse
	var mozambiqueCountry models.Country
	mozambiqueCountry.Name = "Mozambique"
	mozambiqueCountry.Code = "+258"
	mozambiqueCustomers.Country = mozambiqueCountry

	var ugandaCustomers models.CustomersResponse
	var ugandaCountry models.Country
	ugandaCountry.Name = "Uganda"
	ugandaCountry.Code = "+256"
	ugandaCustomers.Country = ugandaCountry

	var noneCustomers models.CustomersResponse
	var noCountry models.Country
	noCountry.Name = "Wrong Numbers"
	noCountry.Code = "--"
	noneCustomers.Country = noCountry

	for i := 0; i < len(customers); i++ {

		var currentCustomer = customers[i]
		isCameron, _ := regexp.MatchString(models.CameronRegex, currentCustomer.Phone)
		isEthiopia, _ := regexp.MatchString(models.EthiopiaRegex, currentCustomer.Phone)
		isMorocco, _ := regexp.MatchString(models.MoroccoRegex, currentCustomer.Phone)
		isMozambique, _ := regexp.MatchString(models.MozambiqueRegex, currentCustomer.Phone)
		isUganda, _ := regexp.MatchString(models.UgandaRegex, currentCustomer.Phone)

		if isCameron {
			fmt.Println(currentCustomer.Name + " from cameron")
			cameronCustomers.Customers = append(cameronCustomers.Customers, currentCustomer)
		} else if isEthiopia {
			fmt.Println(currentCustomer.Name + " from ethiopia")
			ethiopiaCustomers.Customers = append(ethiopiaCustomers.Customers, currentCustomer)
		} else if isMorocco {
			fmt.Println(currentCustomer.Name + " from morocco")
			moroccoCustomers.Customers = append(moroccoCustomers.Customers, currentCustomer)
		} else if isMozambique {
			fmt.Println(currentCustomer.Name + " from mozambique")
			mozambiqueCustomers.Customers = append(cameronCustomers.Customers, currentCustomer)
		} else if isUganda {
			fmt.Println(currentCustomer.Name + " from uganda")
			ugandaCustomers.Customers = append(ugandaCustomers.Customers, currentCustomer)
		} else {
			fmt.Println(currentCustomer.Name + " from none")
			noneCustomers.Customers = append(noneCustomers.Customers, currentCustomer)
		}
	}

	if len(cameronCustomers.Customers) > 0 {
		customersResponse = append(customersResponse, cameronCustomers)
	}

	if len(ethiopiaCustomers.Customers) > 0 {
		customersResponse = append(customersResponse, ethiopiaCustomers)
	}

	if len(moroccoCustomers.Customers) > 0 {
		customersResponse = append(customersResponse, moroccoCustomers)
	}

	if len(mozambiqueCustomers.Customers) > 0 {
		customersResponse = append(customersResponse, mozambiqueCustomers)
	}

	if len(ugandaCustomers.Customers) > 0 {
		customersResponse = append(customersResponse, ugandaCustomers)
	}

	if len(noneCustomers.Customers) > 0 {
		customersResponse = append(customersResponse, noneCustomers)
	}

	return customersResponse
}
