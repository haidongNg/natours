package routers

import (
	v1 "natours/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	// Tours
	apiv1.GET("/tours", v1.GetAllTours)
	apiv1.GET("/tours/:id", v1.GetTour)
	apiv1.POST("/tours", v1.CreateTour)
	apiv1.PATCH("/tours/:id", v1.UpdateTour)
	apiv1.DELETE("/tours/:id", v1.DeleteTour)

	// Users
	apiv1.GET("/users", v1.GetAllUsers)
	apiv1.GET("/users/:id", v1.GetUser)
	apiv1.POST("/users", v1.CreateUser)
	apiv1.PATCH("/users/:id", v1.UpdateUser)
	apiv1.DELETE("/users/:id", v1.DeleteUser)
	return r
}
