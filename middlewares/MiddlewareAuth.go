package middlewares

import "github.com/gin-gonic/gin"

func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//middleware
	}
}
