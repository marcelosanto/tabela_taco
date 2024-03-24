package routes

import (
	"go-alimentos/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	main := router.Group("api/v1")
	{
		alimentos := main.Group("alimentos")
		{
			alimentos.GET("/", controllers.Listar_Alimentos)
			alimentos.POST("/", controllers.Buscar_Alimento)
		}
	}

	return router
}
