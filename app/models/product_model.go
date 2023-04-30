package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	Name    string  `json:"name" bson:"name"`
	Comment string  `json:"comment" bson:"comment"`
	Rating  float64 `json:"rating" bson:"rating"`
}

type Product struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Slug         string             `json:"slug" bson:"slug"`
	Image        string             `json:"image" bson:"image"`
	Images       []string           `json:"images" bson:"images"`
	Brand        string             `json:"brand" bson:"brand"`
	Category     string             `json:"category" bson:"category"`
	Description  string             `json:"description" bson:"description"`
	Price        float64            `json:"price" bson:"price"`
	CountInStock int                `json:"countInStock" bson:"countInStock"`
	Rating       float64            `json:"rating" bson:"rating"`
	NumReviews   int                `json:"numReviews" bson:"numReviews"`
	Reviews      []Review           `json:"reviews" bson:"reviews"`
	CreatedAt    time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
