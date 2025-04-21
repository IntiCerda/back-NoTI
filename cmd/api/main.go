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
	// Obtener la configuración
	config := configs.GetConfig()

	// Establecer el modo de Gin según la configuración
	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Conectar a MongoDB
	mongoClient, err := configs.GetMongoClient()
	if err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(nil); err != nil {
			log.Printf("Error al desconectar de MongoDB: %v", err)
		}
	}()

	// Obtener la base de datos
	mongoDB := configs.GetMongoDB(mongoClient)

	// Inicializar repositorios
	locationRepo := repository.NewLocationRepository(mongoDB)

	// Crear instancia de Resolver con dependencias inyectadas
	resolver := &graph.Resolver{
		LocationRepo: locationRepo,
	}

	// Obtener el schema GraphQL con el resolver
	schema, err := graph.CreateSchema(resolver)
	if err != nil {
		log.Fatalf("Error creando schema GraphQL: %v", err)
	}

	// Configurar handler de GraphQL
	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   config.Debug,
		GraphiQL: config.Debug, // Habilitar GraphiQL solo en modo debug
	})

	// Inicializar router de Gin
	r := gin.Default()

	// Ruta principal para GraphQL
	r.POST("/graphql", handlers.GraphQLHandler(h))

	// Ruta opcional para GraphiQL en modo debug
	if config.Debug {
		r.GET("/graphql", handlers.GraphQLHandler(h))
	}

	// Puerto desde la configuración
	portStr := strconv.Itoa(config.ServerPort)
	serverAddr := fmt.Sprintf(":%s", portStr)

	// Ejecutar servidor
	log.Printf("Servidor iniciado en http://localhost%s/graphql\n", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Error iniciando servidor Gin: %v", err)
	}
}
