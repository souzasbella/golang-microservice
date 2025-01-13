package database

import (
	"context"

	"github.com/souzasbella/golang-microservices/internal/models"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAdress string) ([]models.Customer, error) {
	var customers []models.Customer
	results := c.DB.WithContext(ctx).
		Where(models.Customer{Email: emailAdress}).Find(&customers)
	return customers, results.Error
}
