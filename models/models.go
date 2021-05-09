package models

var (
	CameronRegex    = "\\(237\\)\\ ?[2368]\\d{7,8}$"
	EthiopiaRegex   = "\\(251\\)\\ ?[1-59]\\d{8}$"
	MoroccoRegex    = "\\(212\\)\\ ?[5-9]\\d{8}$"
	MozambiqueRegex = "\\(258\\)\\ ?[28]\\d{7,8}$"
	UgandaRegex     = "\\(256\\)\\ ?\\d{9}$"
)

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Customer struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type CustomersResponse struct {
	Country   Country    `json:"country"`
	Customers []Customer `json:"customers"`
}
