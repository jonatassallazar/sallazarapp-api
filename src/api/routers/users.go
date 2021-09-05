package routers

import (
	"api/src/api/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(g *gin.RouterGroup) {
	usersRoutes := g.Group("/users")
	{
		usersRoutes.GET("/login", controllers.UserLogin)
	}
}
