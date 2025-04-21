package graph

import (
	"errors"
	"fmt"

	"github.com/IntiCerda/gin-graphql-api/internal/models"
	"github.com/IntiCerda/gin-graphql-api/internal/repository"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	LocationRepo *repository.LocationRepository
}

// Simulamos una base de datos con un mapa
var users = map[string]*models.User{
	"1": {ID: "1", Name: "Usuario 1", Email: "usuario1@example.com"},
	"2": {ID: "2", Name: "Usuario 2", Email: "usuario2@example.com"},
}

func (r *Resolver) ResolveUser(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("id debe ser una cadena")
	}

	user, exists := users[id]
	if !exists {
		return nil, errors.New("usuario no encontrado")
	}

	return user, nil
}

func (r *Resolver) ResolveUsers(p graphql.ResolveParams) (interface{}, error) {
	var userList []*models.User
	for _, user := range users {
		userList = append(userList, user)
	}
	return userList, nil
}

func (r *Resolver) CreateUser(p graphql.ResolveParams) (interface{}, error) {
	name, nameOK := p.Args["name"].(string)
	email, emailOK := p.Args["email"].(string)

	if !nameOK || !emailOK {
		return nil, errors.New("argumentos inválidos")
	}

	newID := fmt.Sprintf("%d", len(users)+1)

	newUser := &models.User{
		ID:    newID,
		Name:  name,
		Email: email,
	}

	users[newID] = newUser

	return newUser, nil
}
