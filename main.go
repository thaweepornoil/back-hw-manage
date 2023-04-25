package main

import (
	"back-hw-manage/connections"
	"back-hw-manage/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	// do some initialization here
	fmt.Println("OK")
	connections.Conection_hw()
	connections.Conection()
}

func main() {

	r := gin.Default()
	routeMid := r.Group("/api")
	{
		routeMid.POST("create", controllers.CreateUser)
		routeMid.GET("user/all", controllers.GetUserAll)
	}
	routePublic := r.Group("/")
	{
		routePublic.GET("select", controllers.GetUser)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
