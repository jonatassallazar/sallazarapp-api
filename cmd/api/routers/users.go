package routers

import (
	"api/cmd/api/controllers"
	"api/cmd/api/middlewares"

	"github.com/gin-gonic/gin"
)

// UsersRoutes agrupamento de rotas de endpoints relazionados aos usuários
func UsersRoutes(g *gin.RouterGroup, gc *controllers.GeneralController) {
	usersRoutes := g.Group("/users")
	{
		usersRoutes.Use(middlewares.AuthorizeJWT)
		usersRoutes.GET("/:id", gc.GetUser)
		usersRoutes.PUT("/:id", gc.UpdateUser)
		usersRoutes.DELETE("/:id", gc.DeleteUser)
	}
}
