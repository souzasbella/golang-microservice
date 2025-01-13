package models

type Customer struct {
	CustomerID  string `gorm:"primaryKey" json:"customerId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"emailAdress"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
