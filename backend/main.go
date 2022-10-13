package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joon0000/sa-65/controller"
	"github.com/joon0000/sa-65/entity"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()

	//user route
	r.GET("/users", controller.ListUser)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	//employee route
	r.GET("/employees", controller.ListEmployee)
	r.GET("/employee/:id", controller.GetEmployee)
	r.POST("/employees", controller.CreateEmployee)
	r.PATCH("/employees", controller.UpdateEmployee)
	r.DELETE("/employees/:id", controller.DeleteEmployee)

	//memberClass route
	r.GET("/memberclasses", controller.ListMemberClass)
	r.GET("/memberclass/:id", controller.GetMemberClass)
	r.POST("/memberclasses", controller.CreateMemberClass)
	r.PATCH("/memberclasses", controller.UpdateMemberclass)
	r.DELETE("/memberclasses/:id", controller.DeleteMemberClass)

	//province route
	r.GET("/provinces", controller.ListProvince)
	r.GET("/province/:id", controller.GetProvince)
	r.POST("/provinces", controller.CreateProvince)
	r.PATCH("/provinces", controller.UpdateProvince)
	r.DELETE("/provinces/:id", controller.DeleteMemberClass)

	//role route
	r.GET("/roles", controller.ListRole)
	r.GET("/role/:id", controller.GetRole)
	r.POST("/roles", controller.CreateUser)
	r.PATCH("/roles", controller.UpdateRole)
	r.DELETE("/roles/:id", controller.DeleteRole)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
