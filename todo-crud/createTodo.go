package todocrud

import (
	"context"
	"fmt"
	"time"
    "net/http"

	"github.com/gin-gonic/gin"

	db "banckendproject/connection"
)

func CreateTodo(c *gin.Context) {
	ctx := context.TODO()
	coll := db.ToDoDB.Collection("todos")

	var newTodo ToDoInfo

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	newTodo.CreatedAt = time.Now().Format(time.DateTime)

	fmt.Printf("Document inserted with ID: %s\n", newTodo)
	result, err := coll.InsertOne(ctx, newTodo)

	// result, err := coll.InsertOne(ctx, bson.D{{Key: "title",Value: newTodo.Title},{Key: "description",Value: newTodo.Description},{Key: "created_at",Value: newTodo.CreatedAt}})
	if err != nil {
		panic(err)
	}
	// end insertOne

	// Prints the ID of the inserted document
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	c.IndentedJSON(http.StatusCreated, result)}
