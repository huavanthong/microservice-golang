package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
)

type DiscountService interface {
	GetDiscount(ID string) (*models.GetDiscountResponse, error)
	CreateDiscount(discount *models.Discount) error
	UpdateDiscount(discount *models.Discount) error
	DeleteDiscount(ID string) error
}