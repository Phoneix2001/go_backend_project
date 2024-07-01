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
type SignedDetails struct {
    Number      *int64
    UserName string
    UserId  string
    Exp        int64
	
}

func CreateToken(userData md.UserInfo) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userData.UserName,
			"user_id":  userData.UserId,
			"phone":    userData.Number,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
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

func ValidateToken(ctx *gin.Context) error {
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {

		erMessage.WriteError(ctx, "Missing authorization header")
		return fmt.Errorf("missing authorization header")
	}
	tokenString = tokenString[len("Bearer "):]

	// err := VerifyToken(tokenString)
	// if err != nil {
	// 	erMessage.WriteError(ctx, "Invalid token")
	// 	return fmt.Errorf("invalid token")
	// }
	claims, message := TokenValidator(tokenString)
	if message != "" {
		erMessage.WriteError(ctx,message)
		
		return fmt.Errorf(message);
	}
	ctx.Set("user_id", claims.UserId)
	ctx.Set("username", claims.UserName)
	ctx.Set("number", claims.Number)
	ctx.Set("exp", claims.Exp)
	ctx.Next()
	return nil
}

func TokenValidator(signedToken string) (*SignedDetails,string) {
    token, err := jwt.Parse(
        signedToken,
        func(token *jwt.Token) (interface{}, error) {
            return secretKey, nil
        },
    )
 
    if err != nil {
        msg := err.Error()
        return nil,msg
    }
var claims jwt.MapClaims;
var ok bool;
	 claims, ok = token.Claims.(jwt.MapClaims);
	 var signedDetails SignedDetails;
	 if ok && token.Valid {
		signedDetails.Number = claims["phone"].(*int64)
		signedDetails.UserId = claims["user_id"].(string)
		signedDetails.UserName = claims["username"].(string)
		signedDetails.Exp = claims["exp"].(int64)
	} 


  return &signedDetails,"unable to extract claims" 
}
