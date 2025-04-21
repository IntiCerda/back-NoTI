// internal/graph/resolvers.go
package graph

import (
	"errors"
	"fmt"

	"github.com/IntiCerda/gin-graphql-api/internal/models"
	"github.com/graphql-go/graphql"
)

// Simulamos una base de datos con un mapa
var users = map[string]*models.User{
	"1": {ID: "1", Name: "Usuario 1", Email: "usuario1@example.com"},
	"2": {ID: "2", Name: "Usuario 2", Email: "usuario2@example.com"},
}

// resolveUser busca un usuario por ID
func resolveUser(p graphql.ResolveParams) (interface{}, error) {
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

// resolveUsers devuelve todos los usuarios
func resolveUsers(p graphql.ResolveParams) (interface{}, error) {
	var userList []*models.User
	for _, user := range users {
		userList = append(userList, user)
	}
	return userList, nil
}

// createUser crea un nuevo usuario
func createUser(p graphql.ResolveParams) (interface{}, error) {
	name, nameOK := p.Args["name"].(string)
	email, emailOK := p.Args["email"].(string)

	if !nameOK || !emailOK {
		return nil, errors.New("argumentos inválidos")
	}

	// Generar un ID simple
	newID := fmt.Sprintf("%d", len(users)+1)

	// Crear nuevo usuario
	newUser := &models.User{
		ID:    newID,
		Name:  name,
		Email: email,
	}

	// Almacenar usuario
	users[newID] = newUser

	return newUser, nil
}
