package router

import (
	"github.com/gin-gonic/gin"
	"golang-final-project2-team2/controllers/comment_controllers"
	"golang-final-project2-team2/controllers/photo_controllers"
	"golang-final-project2-team2/controllers/social_media_controllers"
	"golang-final-project2-team2/controllers/user_controllers"
	"golang-final-project2-team2/db"
	"golang-final-project2-team2/middlewares"
	"log"
)

const PORT = ":8080"

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
			photoRouter.GET("/", photo_controllers.GetPhotos)
			photoRouter.PUT("/:photoId", photo_controllers.UpdatePhoto)
			photoRouter.DELETE("/:photoId", photo_controllers.DeletePhoto)
		}
		//router.Use(middlewares.MiddlewareAuth())

		commentRouter := apiRouter.Group("/comments")
		{
			commentRouter.Use(middlewares.MiddlewareAuth())
			commentRouter.POST("/", comment_controllers.CreateComment)
			commentRouter.GET("/", comment_controllers.GetComments)
			commentRouter.PUT("/:commentId", comment_controllers.UpdateComment)
			commentRouter.DELETE("/:commentId", comment_controllers.DeleteComment)
		}

		socialMediaRouter := apiRouter.Group("/social-media")
		{
			socialMediaRouter.Use(middlewares.MiddlewareAuth())
			socialMediaRouter.POST("/", social_media_controllers.CreateSocialMedia)
			socialMediaRouter.GET("/", social_media_controllers.GetSocialMedias)
			socialMediaRouter.PUT("/:socialMediaId", social_media_controllers.UpdateSocialMedia)
			socialMediaRouter.DELETE("/:socialMediaId", social_media_controllers.DeleteSocialMedia)
		}

	}

	err := router.Run(PORT)
	if err != nil {
		log.Fatal(err.Error())
	}
}
