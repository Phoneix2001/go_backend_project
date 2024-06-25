package todocrud

import (
	
	cn "banckendproject/connection"
	"context"
	"fmt"
	"log"
	"net/http"
	
    "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllTODO(c *gin.Context) {

	ctx := context.TODO()
	coll := cn.ToDoDB.Collection("todos")
	//title := "Back to the Future"
	//var todoArray []bson.M
	
	var result []bson.M
	cursor, err := coll.Find(ctx, bson.M{})
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title")
		return
	}

	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var item bson.M

		if err = cursor.Decode(&item); err != nil {
			log.Fatal(err)
		}
		
		result = append(result, item)
		fmt.Println(item)

	}

	fmt.Println("No document was found with the title")
	c.IndentedJSON(http.StatusOK, result)
}

func GetToDoById(c *gin.Context) {
	id := c.Param("id")
	ctx := context.TODO()
	var result []bson.M
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	// fmt.Printf(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	coll := cn.ToDoDB.Collection("todos")
	cursor, err := coll.Find(ctx, filter)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id")
		c.IndentedJSON(http.StatusNotFound, ErrorMessage{Message: "No data found"})
		return
	}

	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item bson.M
		if err = cursor.Decode(&item); err != nil {
			log.Fatal(err)
		}
		result = append(result, item)
		fmt.Println(item)
	}
	if len(result) == 0 {
		c.IndentedJSON(http.StatusNotFound, ErrorMessage{Message: "No data found"})
		return
	}
	// fmt.Println("No document was found with the title")
	c.IndentedJSON(http.StatusOK, result)
}
