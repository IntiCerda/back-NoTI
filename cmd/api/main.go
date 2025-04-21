package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/IntiCerda/gin-graphql-api/configs"
	"github.com/IntiCerda/gin-graphql-api/internal/graph"
	"github.com/IntiCerda/gin-graphql-api/internal/handlers"
	"github.com/IntiCerda/gin-graphql-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}
}

func main() {
	config := configs.GetConfig()

	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	mongoClient, err := configs.GetMongoClient()
	if err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(nil); err != nil {
			log.Printf("Error al desconectar de MongoDB: %v", err)
		}
	}()

	mongoDB := configs.GetMongoDB(mongoClient)

	locationRepo := repository.NewLocationRepository(mongoDB)

	resolver := &graph.Resolver{
		LocationRepo: locationRepo,
	}

	schema, err := graph.CreateSchema(resolver)
	if err != nil {
		log.Fatalf("Error creando schema GraphQL: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   config.Debug,
		GraphiQL: config.Debug,
	})

	r := gin.Default()

	r.POST("/graphql", handlers.GraphQLHandler(h))

	if config.Debug {
		r.GET("/graphql", handlers.GraphQLHandler(h))
	}

	portStr := strconv.Itoa(config.ServerPort)
	serverAddr := fmt.Sprintf(":%s", portStr)

	log.Printf("Servidor iniciado en http://localhost%s/graphql\n", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Error iniciando servidor Gin: %v", err)
	}
}
