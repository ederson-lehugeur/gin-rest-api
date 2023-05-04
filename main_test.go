package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ederson-lehugeur/gin-rest-api/controllers"
	"github.com/ederson-lehugeur/gin-rest-api/database"
	"github.com/ederson-lehugeur/gin-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestGetStudentsSuccessfully(test *testing.T) {
	// TODO - Test with GoMock
	database.Connect()

	routes := SetupRoutes()
	routes.GET("/students", controllers.GetStudents)

	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	routes.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}

func TestCreateStudentWithFailureInValidationCpf(test *testing.T) {
	// TODO - Test with GoMock
	database.Connect()

	routes := SetupRoutes()
	routes.POST("/students", controllers.CreateStudent)

	student := models.Student{Name: "Fulano", CPF: "0000000001", RG: "000000000"}

	requestBody, _ := json.Marshal(student)
	request, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(requestBody))
	response := httptest.NewRecorder()
	routes.ServeHTTP(response, request)

	expectedResponse := `{"error":"CPF: invalid length"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(test, expectedResponse, string(responseBody))
}
