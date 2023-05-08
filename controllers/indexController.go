package controllers

import (
	"net/http"

	"github.com/ederson-lehugeur/gin-rest-api/database"
	"github.com/ederson-lehugeur/gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "notFound.html", nil)
}
