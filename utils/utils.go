package utils

import (
	"encoding/json"
	"fmt"
	

	"golang.org/x/crypto/bcrypt"
)

func PrintStruct(v any)  {
value ,err :=	json.Marshal(v);
if err != nil {
	panic(err)
}
fmt.Print(string(value)) 
}

func Encrypt(pass string) (string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash);
}
 // Hashing the password
 

