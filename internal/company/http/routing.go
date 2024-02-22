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
    companies := r.Group("/companies")
    {
        companies.GET("/:companyid", middleware.CommonHeaders(t.GetByID))
        companies.GET("/", middleware.CommonHeaders(t.List))
        companies.POST("/", middleware.IsAuthorizedJWT(t.Create, "pet-details"))
        companies.PUT("/", middleware.IsAuthorizedJWT(t.Update, "pet-details"))
        companies.DELETE("/:companyid", middleware.IsAuthorizedJWT(t.Delete, "pet-details"))
    }
}
