package middlewares

import (
	"api/cmd/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(c *gin.Context) {
	err := utils.ValidateToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.Next()
}
