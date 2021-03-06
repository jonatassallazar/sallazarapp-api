package routers

import (
	"api/cmd/api/controllers"

	"github.com/gin-gonic/gin"
)

// InitRoutes organiza a estrutura das rotas e o uso de Middlewares
func InitRoutes(r *gin.Engine, gc *controllers.GeneralController) {

	baseApi := r.Group("/v1")
	{
		//Rotas Públicas
		baseApi.POST("/login", gc.UserLogin)
		baseApi.POST("/signup", gc.UserSignup)

		//--------------//

		// Rotas Privadas
		UsersRoutes(baseApi, gc)
		ClientsRoutes(baseApi, gc)
	}
}
