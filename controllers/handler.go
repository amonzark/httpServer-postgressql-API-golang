package controllers

import (
	"httpServer-postgressql-API-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListstudentsHandler(c *gin.Context) {
	student := models.ListStudentsHandler()
	if student == nil || len(student) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, student)
	}
}

func CreatestudentsHandler(c *gin.Context) {
	var std models.Students

	if err := c.BindJSON(&std); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.CreateStudentsHandler(std)
		c.IndentedJSON(http.StatusCreated, std)
	}
}

func Getstudentsbyid(c *gin.Context) {
	id := c.Param("id")

	std := models.GetStudentsbyID(id)

	if std == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, std)
	}
}

func Getstudentsbyage(c *gin.Context) {
	age := c.Param("age")

	std := models.GetStudentsbyAge(age)

	if std == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, std)
	}
}

func Updatestudentsbyid(c *gin.Context) {
	var std models.Students

	if err := c.BindJSON(&std); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.UpdateStudentsbyID(std)
		c.IndentedJSON(http.StatusCreated, std)
	}
}

func Deletestudentsbyid(c *gin.Context) {
	id := c.Param("id")

	std := models.DeleteStudentsbyID(id)

	if std == nil {
		c.IndentedJSON(http.StatusOK, std)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
