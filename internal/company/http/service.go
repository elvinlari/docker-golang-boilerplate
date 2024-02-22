package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/elvinlari/docker-golang/internal/company/domain"
)

type Service struct {
	Service domain.Service
}


func (t *Service) GetByID(c *gin.Context) {
	id := c.Param("companyid")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "companyid must be an integer"})
		return
	}
	company, err := t.Service.GetByID(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("company with id %d not found", idInt)})
		return
	}
	c.IndentedJSON(http.StatusOK, company)
}


func (t *Service) List(c *gin.Context) {
	companies, err := t.Service.List()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, companies)
}


func (t *Service) Create(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company, err := t.Service.Create(request.Company.httpToModel())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, company)
}


func (t *Service) Update(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company, err := t.Service.Update(request.Company.httpToModel())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, company)
}


func (t *Service) Delete(c *gin.Context) {
	id := c.Param("companyid")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "companyid must be an integer"})
		return
	}
	err = t.Service.Delete(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("company with id %d not found", idInt)})
		return
	}
	c.IndentedJSON(http.StatusOK, struct{}{})
}
