package routers

import (
	"api/cmd/api/controllers"
	"api/cmd/api/middlewares"

	"github.com/gin-gonic/gin"
)

// ClientsRoutes agrupamento de rotas de endpoints relacionados aos clientes
func ClientsRoutes(g *gin.RouterGroup, gc *controllers.GeneralController) {
	clientsRoutes := g.Group("/clients")
	{
		clientsRoutes.Use(middlewares.AuthorizeJWT)
		clientsRoutes.GET("/", gc.GetAllClients)
		clientsRoutes.GET("/:id", gc.GetOneClient)
		clientsRoutes.POST("/", gc.AddClient)
		clientsRoutes.PUT("/:id", gc.UpdateClient)
		clientsRoutes.DELETE("/:id", gc.DeleteClient)
	}
}
