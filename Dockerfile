# Usa una imagen base de Go
FROM golang:1.24.2-alpine3.21

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Descarga las dependencias del proyecto
RUN go mod tidy

# Compila la aplicación
RUN go build -o main cmd/api/main.go

# Expone el puerto en el que se ejecutará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
# CMD ["go", "run", "cmd/api/main.go"]
#docker-compose up --build