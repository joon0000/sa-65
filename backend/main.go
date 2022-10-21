package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joon0000/sa-65/controller"
	"github.com/joon0000/sa-65/entity"
	"github.com/joon0000/sa-65/middlewares"
)

//const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			//user routes
			router.GET("/users", controller.ListUser)
			router.GET("/user/:id", controller.GetUser)
			router.POST("/users", controller.CreateUser)
			router.PATCH("/users", controller.UpdateUser)
			router.DELETE("/users/:id", controller.DeleteUser)

			//employee routes
			router.GET("/employees", controller.ListEmployee)
			router.GET("/employee/:id", controller.GetEmployee)
			//r.POST("/employees", controller.CreateEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)
			router.DELETE("/employees/:id", controller.DeleteEmployee)

			//memberClass routes
			router.GET("/memberclasses", controller.ListMemberClass)
			router.GET("/memberclass/:id", controller.GetMemberClass)
			router.POST("/memberclasses", controller.CreateMemberClass)
			router.PATCH("/memberclasses", controller.UpdateMemberclass)
			router.DELETE("/memberclasses/:id", controller.DeleteMemberClass)

			//province routes
			router.GET("/provinces", controller.ListProvince)
			router.GET("/province/:id", controller.GetProvince)
			router.POST("/provinces", controller.CreateProvince)
			router.PATCH("/provinces", controller.UpdateProvince)
			router.DELETE("/provinces/:id", controller.DeleteMemberClass)

			//role routes
			router.GET("/roles", controller.ListRole)
			router.GET("/role/:id", controller.GetRole)
			router.POST("/roles", controller.CreateUser)
			router.PATCH("/roles", controller.UpdateRole)
			router.DELETE("/roles/:id", controller.DeleteRole)

		}
	}

	//Signup User Route
	r.POST("/signup", controller.CreateEmployee)
	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("0.0.0.0:8080")
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
