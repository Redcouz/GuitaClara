package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes Más adelante vamos a inyectar servicios (auth.Service, etc.)
// Por ahora empezamos simple.
func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", healthHandler)

	// Rutas que en el futuro van a ir protegidas
	// api := r.Group("/api")
	// api.Use(AuthMiddleware())
	// {
	// 	   api.GET("/me", meHandler)
	// }
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
