package routes

import (
	"log"

	"github.com/gin-gonic/gin"

	studentController "github.com/ederson-lehugeur/gin-rest-api/controllers"
)

func HandleRequest() {
	r := gin.Default()
	studentsGroup := r.Group("students")
	{
		studentsGroup.GET("", studentController.GetStudents)
		studentsGroup.GET("/:id", studentController.GetStudentById)
		studentsGroup.POST("", studentController.CreateStudent)
		studentsGroup.PATCH("/:id", studentController.EditStudent)
		studentsGroup.DELETE("/:id", studentController.DeleteStudent)
	}
	log.Fatal(r.Run())
}
