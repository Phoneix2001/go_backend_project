package main

import (
	ar   "banckendproject/auth/router"
	cn   "banckendproject/connection"
	todo "banckendproject/todo-crud"
    "github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	cn.Mongodb()

	ar.AuthInit(router)
    todo.TodoInit(router)
	router.Run("0.0.0.0")
}