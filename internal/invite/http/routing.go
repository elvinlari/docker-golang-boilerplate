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
    invites := r.Group("/invites")
    {
        invites.GET("/:inviteid", middleware.CommonHeaders(t.GetByID))
        invites.GET("/", middleware.CommonHeaders(t.List))
        invites.POST("/", middleware.IsAuthorizedJWT(t.Create, "pet-details"))
        invites.PUT("/", middleware.IsAuthorizedJWT(t.Update, "pet-details"))
        invites.DELETE("/:inviteid", middleware.IsAuthorizedJWT(t.Delete, "pet-details"))
    }
}
