// internal/handlers/graphql.go
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

// GraphQLHandler adapta el handler de graphql-go para usarlo con Gin
func GraphQLHandler(h *handler.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
