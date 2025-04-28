package repository

import (
	"context"
	"time"

	"github.com/IntiCerda/gin-graphql-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationRepository struct {
	collection *mongo.Collection
}

func NewLocationRepository(db *mongo.Database) *LocationRepository {
	return &LocationRepository{
		collection: db.Collection("locations"),
	}
}

func (r *LocationRepository) InsertLocation(ctx context.Context, loc *models.Location) (primitive.ObjectID, error) {
	loc.CreatedAt = time.Now()
	result, err := r.collection.InsertOne(ctx, loc)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *LocationRepository) GetAllLocations(ctx context.Context) ([]*models.Location, error) {
	var locations []*models.Location
	cutoff := time.Now().Add(-6 * time.Hour)
	filter := bson.M{"createdAt": bson.M{"$gte": cutoff}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var loc models.Location
		if err := cursor.Decode(&loc); err != nil {
			return nil, err
		}
		locations = append(locations, &loc)
	}
	return locations, nil
}

func (r *LocationRepository) GetLocationByID(ctx context.Context, id primitive.ObjectID) (*models.Location, error) {
	var loc models.Location
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&loc)
	if err != nil {
		return nil, err
	}
	return &loc, nil
}
