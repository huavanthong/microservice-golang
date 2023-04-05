package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/models"
)

type BasketService interface {
	CreateBasket(cbr *models.CreateBasketRequest) (*entities.Basket, error)
	GetBasket(userName string) (*entities.Basket, error)
	UpdateBasket(userName string, cart *entities.Basket) (*entities.Basket, error)
	DeleteBasket(userName string) error
}
