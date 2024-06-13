package main

import (
	todo "banckendproject/todo-crud"

	cn "banckendproject/connection"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cn.Mongodb()



	
	router.GET("/todos", func(ctx *gin.Context) {
		todo.GetAllTODO(ctx)
	})
	router.POST("/todos", func(ctx *gin.Context) {
		todo.CreateTodo(ctx)
	})
	router.GET("/todos/:id", func(ctx *gin.Context) {
		todo.GetToDoById(ctx)
	})

	router.PUT("/todos/:id",func(ctx *gin.Context) {
		
	})
	router.Run("localhost:8080")
}

// func getAlbums(c *gin.Context, jsonData bson.M) {
// 	c.IndentedJSON(http.StatusOK, jsonData)
// }

// func createAlbum(c *gin.Context) {
// 	var newAlbum album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	newAlbum.ID = strconv.Itoa(len(albums)+1)
// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found", "status": 404})
// }
