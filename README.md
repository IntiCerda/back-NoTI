# API GraphQL con Gin en Go

Este proyecto es una API GraphQL implementada en Go utilizando el framework Gin y la biblioteca graphql-go.

## Requisitos

- Go 1.16 o superior
- Git

## Instalación

1. Clonar el repositorio
```
git clone https://github.com/tuusuario/gin-graphql-api.git
cd gin-graphql-api
```

2. Instalar dependencias
```
go mod tidy
```

3. Ejecutar el servidor
```
go run cmd/api/main.go
```

El servidor estará disponible en http://localhost:8080/graphql

## Ejemplos de queries

### Obtener un usuario
```graphql
{
  user(id: "1") {
    id
    name
    email
  }
}
```

### Obtener todos los usuarios
```graphql
{
  users {
    id
    name
    email
  }
}
```

### Crear un usuario
```graphql
mutation {
  createUser(name: "Nuevo Usuario", email: "nuevo@example.com") {
    id
    name
    email
  }
}
```s