package user

import (
	"net/http"

	"mux-server/dto"
	"mux-server/utils"

	"go.uber.org/zap"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	zap.L().Debug("GetUserDetails")

	//gets value from the fuction and stores it in the struct
	userDetails, err := getUserInfo()
	if err != nil {
		zap.L().Error("GetUserDetails", zap.Error(err))
		res := &dto.UserDetailsRes{
			Code: http.StatusInternalServerError,
		}
		utils.Response(w, res)
		return
	}

	//takes value from the struct and updates in response variable
	res := &dto.UserDetailsRes{
		Code:    http.StatusOK,
		Message: userDetails.Message,
		Name:    userDetails.Name,
		Country: userDetails.Country,
	}

	//sends a response
	utils.Response(w, res)
}

//adds user info to the struct
func getUserInfo() (*dto.UserDetailsRes, error) {
	userInfo := &dto.UserDetailsRes{
		Code:     200,
		Message:  "Successfull",
		Name:     "Saurabh",
		Country:  "India",
		Services: "None",
	}
	return userInfo, nil
}
