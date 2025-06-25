package routers

import (
	"dashboard-backend/controllers"

	"github.com/gin-gonic/gin"
)

func contactRouter(r *gin.RouterGroup) {
	r.GET("/", controllers.GetContacts)
	r.GET("/:id", controllers.GetContactByID)
	r.POST("/", controllers.AddContact)
	r.PATCH("/:id", controllers.UpdateContact)
	r.DELETE("/:id", controllers.DeleteContact)
}
