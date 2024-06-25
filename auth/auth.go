package auth

import (
	md "banckendproject/auth/model"
	vl "banckendproject/auth/validator"
	db "banckendproject/connection"
	dao "banckendproject/dao"
	ut "banckendproject/utils"
	erMessage "banckendproject/utils/error"
	er "banckendproject/utils/model"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// type GetProfileProps struct {
// 	Ctx *gin.Context
//     Id string
// }

func SignUp(ctx *gin.Context) {
	var userData md.UserInfo

	if err := ctx.BindJSON(&userData); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	var isValidated = vl.ValidateSignUp(ctx, userData)
	if isValidated {
		nowTime := time.Now().Format(time.DateTime)
		userData.CreatedAt = nowTime
		userData.UpdatedAt = nowTime
		var id = primitive.NewObjectID()
		userData.UserId = &id
		userData.IsDeleted = false
		c := context.TODO()
		coll := db.UsersDB.Collection("user_details")
		additinalColl := db.UsersDB.Collection("additional_details")

		ut.PrintStruct(userData)
		userDD := userData.DeviceDetails
		userData.DeviceDetails = nil
		userData.Password = ut.Encrypt(userData.Password)
		if accessTokenString, err := ut.CreateToken(userData); err == nil {
			userData.AccessToken = accessTokenString
		} else if err != nil {
			panic(err)
		}
		result, err := coll.InsertOne(c, userData)
		if err != nil {
			erMessage.WriteError(ctx, "Username already exist")
			return
		}

		userDD.UserId = userData.UserId

		if _, err := additinalColl.InsertOne(c, userDD); err != nil {
			erMessage.WriteError(ctx, err.Error())
			return
		}

		// end insertOne

		// Prints the ID of the inserted document
		fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
		userData.DeviceDetails = userDD
		userData.Password = ""
		userData.CreatedAt = ""
		userData.UpdatedAt = ""
		ctx.IndentedJSON(http.StatusCreated, userData)
	}

}

func SignIn(ctx *gin.Context) {
	var lgCreds md.LoginCreds
	var userData md.UserInfo

	if err := ctx.BindJSON(&lgCreds); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	var isValidated = vl.ValidateSignLogin(ctx, lgCreds)
	if isValidated {
		c := context.TODO()
		coll := db.UsersDB.Collection("user_details")
		// key ,errors :=	primitive.ObjectIDFromHex(lgCreds.Username);
		// if errors != nil {
		// 	panic(errors)
		// }
		cursor, err := coll.Aggregate(c, dao.UserDataAG(lgCreds.Username))
		if err != nil {
			panic(err)
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			if err = cursor.Decode(&userData); err != nil {
				log.Fatal(err)
			}

		}
		fmt.Printf("User data retrived")
		fmt.Print(userData)
		if userData.UserName == "" {
			erMessage.WriteError(ctx, "Username not found")

			return
		}
		if bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(lgCreds.Password)) != nil {
			erMessage.WriteError(ctx, "Password is not correct")
			return
		}
		if accessToken, err := ut.CreateToken(userData); err == nil {
			userData.AccessToken = accessToken
		} else if err != nil {
			panic(err)
		}
		coll.UpdateOne(c, bson.D{{Key: "user_id", Value: userData.UserId}}, dao.UpdateUser(userData.AccessToken))
		userData.Password = ""
		ctx.IndentedJSON(http.StatusOK, userData)
	}

}

func GetProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	var result []bson.M

	coll := db.UsersDB.Collection("user_details")
	cursor, err := coll.Aggregate(ctx, dao.UserDataByUserId(id))
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id")
		erMessage.WriteError(ctx, "No data found")
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

	}
	if len(result) == 0 {
		erMessage.WriteError(ctx, "No data found")
		return
	}
	// fmt.Println("No document was found with the title")
	ctx.IndentedJSON(http.StatusOK, result)
}
func LogOut(ctx *gin.Context) {

	accessToken := ut.ValidateToken(ctx)
	if accessToken == "" {
		return
	}

	c := context.TODO()
	coll := db.UsersDB.Collection("user_details")
	coll.UpdateOne(c, bson.D{{Key: "user_id", Value: accessToken}},dao.UpdateUser(accessToken))

}
