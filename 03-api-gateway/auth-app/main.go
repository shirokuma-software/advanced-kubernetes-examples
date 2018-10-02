package main

import (
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("x-auth-token")

		if authToken != "cute-shirokuma" {
			c.String(401, "Failed to authentication")
			c.Abort()
			return
		}
	}
}

func GetRoot(c *gin.Context) {
	c.String(200, "ok")
}

func Router(route *gin.Engine) {
	authorized := route.Group("/")

	authorized.Use(Authorization())
	{
		authorized.GET("/", GetRoot)
	}
}

func main() {
	app := gin.Default()
	Router(app)
	app.Run(":8000")
}
