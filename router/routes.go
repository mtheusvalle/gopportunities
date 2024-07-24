package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mtheusvalle/gopportunities/docs"
	"github.com/mtheusvalle/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwager "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
	// Initialize swagger
	router.GET("/swagger/*any", ginSwager.WrapHandler(swaggerfiles.Handler))
}
