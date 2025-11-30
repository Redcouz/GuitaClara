package httpapi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware es el punto donde vamos a:
// - leer el header Authorization
// - validar el token (Cognito / IdentityProvider)
// - poner los claims en el contexto
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
			return
		}

		token := parts[1]

		// TODO: acá vamos a llamar al IdentityProvider
		// y setear, por ejemplo:
		//
		//  c.Set("user_email", claims.Email)
		//  c.Set("user_sub", claims.Sub)
		//
		// Por ahora, lo dejamos pasar para compilar
		_ = token

		c.Next()
	}
}
