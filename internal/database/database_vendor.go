package database

import (
	"context"

	"github.com/sanjevscet/go-microservices/internal/models"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor

	result := c.DB.WithContext(ctx).Where(models.Service{}).Find(&vendors)

	return vendors, result.Error
}
