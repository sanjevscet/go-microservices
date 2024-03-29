package database

import (
	"context"

	"github.com/sanjevscet/go-microservices/internal/models"
)

func (c Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
	var products []models.Product

	result := c.DB.WithContext(ctx).Where(models.Product{VendorID: vendorId}).Find(&products)

	return products, result.Error
}
