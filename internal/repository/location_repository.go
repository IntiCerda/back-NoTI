// internal/repository/location_repository.go
package repository

import (
	"context"
	"time"

	"github.com/IntiCerda/gin-graphql-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const locationCollection = "WazePeruvian"

// LocationRepository maneja las operaciones de base de datos para locations
type LocationRepository struct {
	collection *mongo.Collection
}

// NewLocationRepository crea un nuevo repositorio de locations
func NewLocationRepository(db *mongo.Database) *LocationRepository {
	collection := db.Collection(locationCollection)

	// Crear índice geoespacial
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "location", Value: "2dsphere"},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		// En producción, manejar este error adecuadamente
		// Para este ejemplo, simplemente lo imprimimos
		println("Error creando índice geoespacial:", err.Error())
	}

	return &LocationRepository{collection: collection}
}

// CreateLocation inserta una nueva ubicación en la base de datos
func (r *LocationRepository) CreateLocation(lat, lng float64, comment string) (*models.Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()

	location := models.Location{
		Latitude:  lat,
		Longitude: lng,
		Comment:   comment,
		CreatedAt: now,
		Location: struct {
			Type        string    `json:"type" bson:"type"`
			Coordinates []float64 `json:"coordinates" bson:"coordinates"`
		}{
			Type:        "Point",
			Coordinates: []float64{lng, lat}, // MongoDB usa [longitude, latitude]
		},
	}

	result, err := r.collection.InsertOne(ctx, location)
	if err != nil {
		return nil, err
	}

	// Obtener el ID generado y asignarlo a la ubicación
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		location.ID = oid
	}

	return &location, nil
}

// GetLocationByID obtiene una ubicación por su ID
func (r *LocationRepository) GetLocationByID(id string) (*models.Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var location models.Location
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&location)
	if err != nil {
		return nil, err
	}

	return &location, nil
}

// GetLocations obtiene todas las ubicaciones
func (r *LocationRepository) GetLocations() ([]*models.Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ordenar por fecha de creación descendente
	findOptions := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var locations []*models.Location
	if err = cursor.All(ctx, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}

// GetNearbyLocations obtiene ubicaciones cercanas a un punto
func (r *LocationRepository) GetNearbyLocations(lat, lng, radiusKm float64) ([]*models.Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Consulta geoespacial: encontrar puntos dentro de un radio
	// $nearSphere busca puntos cercanos a una coordenada específica
	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lng, lat}, // MongoDB usa [longitude, latitude]
				},
				"$maxDistance": radiusKm * 1000, // Convertir a metros
			},
		},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var locations []*models.Location
	if err = cursor.All(ctx, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}
