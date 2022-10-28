package router

import (
	"golang-final-project2-team2/controllers/UserController"
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

const PORT = ":8081"

func init() {
	db.InitializeDB()
}

func StartRouter() {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", UserController.CreateUser)
		userRouter.POST("/login", UserController.CreateUser)
		userRouter.Use(middlewares.MiddlewareAuth())
		userRouter.PUT("/users", UserController.CreateUser)

		//productRouter.PUT("/:productId", UserController.UpdateProduct)
		//productRouter.GET("/", UserController.GetProducts)
		//productRouter.DELETE("/:productId", UserController.DeleteProduct)
	}
	//router.Use(middlewares.MiddlewareAuth())

	err := router.Run(PORT)
	if err != nil {
		log.Fatal(err.Error())
	}
}
