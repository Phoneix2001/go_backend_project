package todocrud

import "github.com/gin-gonic/gin"

func TodoInit(router *gin.Engine) {
	rt := router.Group("/user")
	{
		rt.GET("/todos", GetAllTODO)

		rt.POST("/todos",CreateTodo)
		rt.GET("/todos/:id",GetToDoById)
rt.PUT("/todos/:id", func(ctx *gin.Context) {

		})
	}
}
