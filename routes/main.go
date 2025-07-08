package routes

import (
	docs "dashboard-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CombineRouter(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	contactRouter(r.Group("/contacts"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
