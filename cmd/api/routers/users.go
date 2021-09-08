package routers

import (
	"api/cmd/api/controllers"
	"api/cmd/api/middlewares"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(g *gin.RouterGroup, gc *controllers.GeneralController) {
	usersRoutes := g.Group("/users")
	{
		usersRoutes.Use(middlewares.AuthorizeJWT)
		usersRoutes.GET("/:id", gc.GetUser)
	}
}
