package routes

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ederson-lehugeur/gin-rest-api/controllers"
	studentController "github.com/ederson-lehugeur/gin-rest-api/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/", controllers.Index)
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	studentsGroup := r.Group("students")
	{
		studentsGroup.GET("", studentController.GetStudents)
		studentsGroup.GET("/:id", studentController.GetStudentById)
		studentsGroup.POST("", studentController.CreateStudent)
		studentsGroup.PATCH("/:id", studentController.EditStudent)
		studentsGroup.DELETE("/:id", studentController.DeleteStudent)
	}
	r.NoRoute(controllers.NotFound)
	log.Fatal(r.Run())
}
