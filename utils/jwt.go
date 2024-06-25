package utils

import (
	md "banckendproject/auth/model"
	erMessage "banckendproject/utils/error"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)
var secretKey = []byte("2i3846923ksjadfk34fd#&")

func CreateToken(userData md.UserInfo) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": userData.UserName,
		"user_id" : userData.UserId,
		"phone"   : userData.Number,
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return secretKey, nil
	})
   
	if err != nil {
	   return err
	}
   
	if !token.Valid {
	   return fmt.Errorf("invalid token")
	}
   
	return nil
 }

 func ValidateToken(ctx *gin.Context) (string) {
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
	 
	  erMessage.WriteError(ctx, "Missing authorization header")
	  return ""
	}
	tokenString = tokenString[len("Bearer "):]
	
	err := VerifyToken(tokenString)
	if err != nil {
	  erMessage.WriteError(ctx, "Invalid token");
	  return ""
	}
	return tokenString;
 }