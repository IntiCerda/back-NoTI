package graph

import (
	"context"
	"errors"
	"time"

	"github.com/IntiCerda/gin-graphql-api/internal/models"
	"github.com/IntiCerda/gin-graphql-api/internal/repository"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resolver struct {
	LocationRepo *repository.LocationRepository
}

func (r *Resolver) ResolveLocations(p graphql.ResolveParams) (interface{}, error) {
	return r.LocationRepo.GetAllLocations(context.Background())
}

func (r *Resolver) ResolveLocationByID(p graphql.ResolveParams) (interface{}, error) {
	idStr, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("id inválido")
	}
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, err
	}
	return r.LocationRepo.GetLocationByID(context.Background(), id)
}

//Funcion para retonar comentario activos (2 horas)

func (r *Resolver) CreateLocation(p graphql.ResolveParams) (interface{}, error) {
	latitude, latOK := p.Args["latitude"].(float64)
	longitude, longOK := p.Args["longitude"].(float64)
	title, _ := p.Args["title"].(string)
	category, _ := p.Args["category"].(string)
	comment, _ := p.Args["comment"].(string)

	if !latOK || !longOK {
		return nil, errors.New("latitud y longitud son requeridas")
	}

	loc := &models.Location{
		Latitude:  latitude,
		Longitude: longitude,
		Comment:   comment,
		Title:     title,
		Category:  category,
		CreatedAt: time.Now(),
	}
	loc.Location.Type = "Point"
	loc.Location.Coordinates = []float64{longitude, latitude}

	inserted, err := r.LocationRepo.InsertLocation(context.Background(), loc)
	if err != nil {
		return nil, err
	}

	loc.ID = inserted
	return loc, nil
}
