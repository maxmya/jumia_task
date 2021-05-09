package models

import (
	"fmt"
	"regexp"
)

var (
	CameronRegex    = "\\(237\\)\\ ?[2368]\\d{7,8}$"
	EthiopiaRegex   = "\\(251\\)\\ ?[1-59]\\d{8}$"
	MoroccoRegex    = "\\(212\\)\\ ?[5-9]\\d{8}$"
	MozambiqueRegex = "\\(258\\)\\ ?[28]\\d{7,8}$"
	UgandaRegex     = "\\(256\\)\\ ?\\d{9}$"
)

type Customer struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type CustomersResponse struct {
	Country   string     `json:"country"`
	Code      string     `json:"code"`
	Customers []Customer `json:"customers"`
}

func CategorizeCustomersByPhoneNumbers(customers []Customer) []CustomersResponse {

	var customersResponse []CustomersResponse

	var cameronCustomers CustomersResponse
	cameronCustomers.Country = "Cameron"
	cameronCustomers.Code = "+237"

	var ethiopiaCustomers CustomersResponse
	ethiopiaCustomers.Country = "Ethiopia"
	ethiopiaCustomers.Code = "+251"

	var moroccoCustomers CustomersResponse
	moroccoCustomers.Country = "Morocco"
	moroccoCustomers.Code = "+212"

	var mozambiqueCustomers CustomersResponse
	mozambiqueCustomers.Country = "Mozambique"
	mozambiqueCustomers.Code = "+258"

	var ugandaCustomers CustomersResponse
	ugandaCustomers.Country = "Uganda"
	ugandaCustomers.Code = "+256"

	var noneCustomers CustomersResponse
	noneCustomers.Country = "None"

	for i := 0; i < len(customers); i++ {

		var currentCustomer = customers[i]
		isCameron, _ := regexp.MatchString(CameronRegex, currentCustomer.Phone)
		isEthiopia, _ := regexp.MatchString(EthiopiaRegex, currentCustomer.Phone)
		isMorocco, _ := regexp.MatchString(MoroccoRegex, currentCustomer.Phone)
		isMozambique, _ := regexp.MatchString(MozambiqueRegex, currentCustomer.Phone)
		isUganda, _ := regexp.MatchString(UgandaRegex, currentCustomer.Phone)

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

	customersResponse = append(customersResponse, cameronCustomers)
	customersResponse = append(customersResponse, ethiopiaCustomers)
	customersResponse = append(customersResponse, moroccoCustomers)
	customersResponse = append(customersResponse, mozambiqueCustomers)
	customersResponse = append(customersResponse, ugandaCustomers)
	customersResponse = append(customersResponse, noneCustomers)

	return customersResponse
}
