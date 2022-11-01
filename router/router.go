package router

import (
	"github.com/gin-gonic/gin"
	"golang-final-project2-team2/controllers/photo_controllers"
	"golang-final-project2-team2/controllers/user_controllers"
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/middlewares"
	"log"
)

const PORT = ":8081"

func init() {
	db.InitializeDB()
}

func StartRouter() {
	router := gin.Default()
	apiRouter := router.Group("/api")
	{
		userRouter := apiRouter.Group("/users")
		{
			userRouter.POST("/register", user_controllers.CreateUser)
			userRouter.POST("/login", user_controllers.UserLogin)
			userRouter.Use(middlewares.MiddlewareAuth())
			userRouter.PUT("/:userId", user_controllers.UpdateUser)
			userRouter.DELETE("/", user_controllers.DeleteUser)
		}

		photoRouter := apiRouter.Group("/photos")
		{
			photoRouter.Use(middlewares.MiddlewareAuth())
			photoRouter.POST("/", photo_controllers.CreatePhoto)
		}
		//router.Use(middlewares.MiddlewareAuth())

	}

	err := router.Run(PORT)
	if err != nil {
		log.Fatal(err.Error())
	}
}
