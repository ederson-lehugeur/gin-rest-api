package controllers

import (
	"net/http"

	"github.com/ederson-lehugeur/gin-rest-api/database"
	"github.com/ederson-lehugeur/gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	studentParameters := models.Student{}

	name, exist := c.GetQuery("name")
	if exist {
		studentParameters.Name = name
	}

	cpf, exist := c.GetQuery("cpf")
	if exist {
		studentParameters.CPF = cpf
	}

	rg, exist := c.GetQuery("rg")
	if exist {
		studentParameters.RG = rg
	}

	database.DB.Where(&studentParameters).Limit(10).Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Save(&student)
	c.JSON(http.StatusCreated, student)
}

func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Deleted Resource",
	})
}
