package validator

import (
	md "banckendproject/auth/model"
	erMessage "banckendproject/utils/error"
	er "banckendproject/utils/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateSignUp(ctx *gin.Context, userData md.UserInfo) bool {
	if strings.TrimSpace(userData.FirstName) == "" {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "first name not allowed to be empty"})
		return false
	}
	if strings.TrimSpace(userData.LastName) == "" {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "last name not allowed to be empty"})
		return false
	}
	if strings.TrimSpace(userData.Age) == "" {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "age not allowed to be empty"})
		return false
	}
	if strings.TrimSpace(userData.UserName) == "" {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "username not allowed to be empty"})
		return false
	}
	if strings.TrimSpace(userData.Mail) == "" {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "mail not allowed to be empty"})
		return false
	}
	if userData.Number == nil {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "number not allowed to be empty"})
		return false
	}
	if userData.DeviceDetails != nil {
		if strings.TrimSpace(userData.DeviceDetails.Modalname) == "" {
			ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "device details not allowed to be empty"})
			return false
		}
		if strings.TrimSpace(userData.DeviceDetails.AppVersion) == "" {
			ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "app version not allowed to be empty"})
			return false
		}
		if strings.TrimSpace(userData.DeviceDetails.DeviceOS) == "" {
			ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "device os not allowed to be empty"})
			return false
		}
		if strings.TrimSpace(userData.DeviceDetails.DeviceType) == "" {
			ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "device type not allowed to be empty"})
			return false
		}
		if userData.DeviceDetails.BatLvl == nil {
			ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "battery level not allowed to be empty"})
			return false
		}
		if strings.TrimSpace(userData.DeviceDetails.StoreBundleID) == "" {
			ctx.IndentedJSON(http.StatusBadGateway, er.ErrorMessage{Status: http.StatusBadGateway, Message: "store bundle id is not allowed to be empty"})
			return false
		}
		if strings.TrimSpace(userData.DeviceDetails.StoreVersion) == "" {
			ctx.IndentedJSON(http.StatusBadGateway, er.ErrorMessage{Status: http.StatusBadRequest, Message: "store version is not allowed to be empty"})
			return false
		}
	}
	if strings.TrimSpace(userData.Gender) == "" {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "Gender not allowed to be empty"})
		return false
	}
	if strings.TrimSpace(userData.Education) == "" {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "Education not allowed to be empty"})
		return false
	}
	if userData.Lat == nil {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "Lat not allowed to be empty"})
		return false
	}
	if userData.Lng == nil {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "Lng not allowed to be empty"})
		return false
	}
	if strings.TrimSpace(userData.Password) == "" {
		ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: "Password not allowed to be empty"})
		return false
	}
	return true
}

func ValidateSignLogin(ctx *gin.Context, lgCreds md.LoginCreds) bool {
	if lgCreds.Username == "" {
		erMessage.WriteError(ctx, "username is not allowed to be empty")
		return false
	}
	if lgCreds.Password == "" {
		erMessage.WriteError(ctx, "Password is not allowed to be empty")
	}
	return true
}


