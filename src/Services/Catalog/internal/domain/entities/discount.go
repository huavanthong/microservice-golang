package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/************************ Define structure product ************************/
type Discount struct {
	ID              primitive.ObjectID `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	DiscountPercent float32            `json:"discountpercent" bson:"discountpercent"`
	Active          bool               `json:"active" bson:"active"`
	CreatedAt       string             `json:"created_at" bson:"created_at"`
	UpdatedAt       string             `json:"updated_at" bson:"updated_at"`
	DeleteAt        string             `json:"deleted_at" bson:"deleted_at"`
}
