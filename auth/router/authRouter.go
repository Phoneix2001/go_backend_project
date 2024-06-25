package authrouter

import (
	au "banckendproject/auth"
	"github.com/gin-gonic/gin"
)

func AuthInit(router *gin.Engine) {
	rt := router.Group("/user")
	{
		rt.GET("/getProfile", au.GetProfile)
		rt.GET("/getProfile/:id", au.GetProfile)

		rt.POST("/signin", au.SignIn)
		rt.POST("/signup", au.SignUp)

		rt.PUT("/logout", au.LogOut)
	}
}
