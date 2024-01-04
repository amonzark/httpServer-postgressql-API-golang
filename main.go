package main

import (
	"httpServer-postgressql-API-golang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/students", controllers.ListstudentsHandler)            //listing
	router.POST("/students/create", controllers.CreatestudentsHandler)  //create
	router.GET("/students/get/:id", controllers.Getstudentsbyid)        //get by id
	router.POST("/students/update/:id", controllers.Updatestudentsbyid) //update by id
	router.POST("/students/delete/:id", controllers.Deletestudentsbyid) //delete by id
	router.GET("/students/getby/:age", controllers.Getstudentsbyage)    //get by age
	router.Run("localhost:8085")
}
