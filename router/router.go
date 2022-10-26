package router

import (
	"golang-final-project2-team2/controller/user_controller"
	"golang-final-project2-team2/db"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func init() {
	db.InitializeDB()
}

func StartRouter() {
	router := gin.Default()

	productRouter := router.Group("/products")
	{
		productRouter.POST("/", user_controller.CreateProduct)
		productRouter.PUT("/:productId", user_controller.UpdateProduct)
		productRouter.GET("/", user_controller.GetProducts)
		productRouter.DELETE("/:productId", user_controller.DeleteProduct)
	}

	router.Run(PORT)
}
