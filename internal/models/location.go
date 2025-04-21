package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Latitude  float64            `json:"latitude" bson:"latitude"`
	Longitude float64            `json:"longitude" bson:"longitude"`
	Comment   string             `json:"comment" bson:"comment"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	// MongoDB usa GeoJSON para coordenadas
	Location struct {
		Type        string    `json:"type" bson:"type"`
		Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	} `json:"location" bson:"location"`
}
