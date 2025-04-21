package main

import (
	"log"

	"github.com/IntiCerda/gin-graphql-api/internal/graph"
	"github.com/IntiCerda/gin-graphql-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main() {
	// Inicializar router de Gin
	r := gin.Default()

	// Obtener el schema GraphQL
	schema, err := graph.CreateSchema()
	if err != nil {
		log.Fatalf("Error creando schema GraphQL: %v", err)
	}

	// Configurar handler de GraphQL
	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Configurar ruta para GraphQL
	r.POST("/graphql", handlers.GraphQLHandler(h))
	r.GET("/graphql", handlers.GraphQLHandler(h)) // Para GraphiQL

	// Ejecutar servidor
	log.Println("Servidor iniciado en http://localhost:8080/graphql")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error iniciando servidor Gin: %v", err)
	}
}
