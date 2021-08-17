package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type student struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Class string  `json:"class"`
  
}


var students = []student{
    {ID: "1", Name: "Ninh", Class: "12A9"},
	{ID: "1", Name: "tien", Class: "12A7"},
}

func main() {
    router := gin.Default()
    router.GET("/students", getStudents)
    router.GET("/students/:id", studentID)
    router.POST("/albums", getStudentID)

    router.Run("localhost:8080")
}

func getStudents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, students)
}


func studentID(c *gin.Context) {
    var newStudent  student

  
    if err := c.BindJSON(&newStudent); err != nil {
        return
    }

   
    students = append(students, newStudent)
    c.IndentedJSON(http.StatusCreated, newStudent)
}


func getStudentID(c *gin.Context) {
    id := c.Param("id")

    
    for _, a := range students {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}