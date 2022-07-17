package services

import "github.com/huavanthong/microservice-golang/user-api-v3/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
	UpsertUser(string, *models.UpdateDBUser) (*models.DBResponse, error)
}