package http

import (
	"github.com/gin-gonic/gin"
	"github.com/elvinlari/docker-golang/cmd/middleware"
)


type ServiceHTTPHandlers interface {
	GetByID(c *gin.Context)
	List(c *gin.Context)
	Create(c *gin.Context)
    Update(c *gin.Context)
	Delete(c *gin.Context)
}

func RegisterRoutes(r *gin.Engine, t ServiceHTTPHandlers) {
    users := r.Group("/users")
    {
        users.GET("/:userid", middleware.CommonHeaders(t.GetByID))
        users.GET("/", middleware.CommonHeaders(t.List))
        users.POST("/", middleware.IsAuthorizedJWT(t.Create, "pet-details"))
        users.PUT("/", middleware.IsAuthorizedJWT(t.Update, "pet-details"))
        users.DELETE("/:userid", middleware.IsAuthorizedJWT(t.Delete, "pet-details"))
    }
}
