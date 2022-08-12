package user

import (
	"net/http"

	"mux-server/dto"
	"mux-server/utils"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func PostService(w http.ResponseWriter, r *http.Request) {
	zap.L().Debug("PostUserDetails")

	//takes the value from the request
	params := mux.Vars(r)
	service := params["service"]

	//setting value to the struct
	setDetails, err := setUserInfo(service)
	if err != nil {
		zap.L().Error("SetUserDetails", zap.Error(err))
		res := &dto.UserDetailsRes{
			Code: http.StatusInternalServerError,
		}
		utils.Response(w, res)
		return
	}

	//takes value from the struct and updates in response variable
	res := &dto.UserDetailsRes{
		Code:     http.StatusOK,
		Message:  setDetails.Message,
		Name:     setDetails.Name,
		Country:  setDetails.Country,
		Services: setDetails.Services,
	}

	//sends a response
	utils.Response(w, res)
}

//adds updated info in the struct
func setUserInfo(service string) (*dto.UserDetailsRes, error) {
	userInfo := &dto.UserDetailsRes{
		Code:     200,
		Message:  "Successfull",
		Name:     "Saurabh",
		Country:  "India",
		Services: service,
	}
	return userInfo, nil
}
