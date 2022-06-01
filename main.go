package main

//gin
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-base/controller"
	"web-base/utils"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/auth/ping", utils.JWTAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	user := r.Group("/user")
	{
		user.POST("/login", controller.Login)
		user.POST("/register", controller.Register)
		user.GET("/info", utils.JWTAuth(), controller.GetUserInfo)
	}
	r.Run() // listen and serve on
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST,GET, OPTIONS, DELETE,PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
