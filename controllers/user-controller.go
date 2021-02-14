package controllers

import (
	"encoding/json"
	base_common "github.com/lambda-0/base-common/base-common"
	"github.com/lambda-0/base-common/data"
	"github.com/lambda-0/base-common/models"
	"net/http"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		base_common.DisplayAppError(w,
			err,
			"Invalid user data",
			500,
		)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	userCol := context.Collection(models.UserCollection)
	userRepo := &data.UserRepository(userCol)
	userRepo.CreateUser(user)
	user.HashPassword = nil
	if resp, err := json.Marshal(UserResource{Data: *user}); err != nil {
		base_common.DisplayAppError(w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(resp)
	}


}

func UserLogin(w http.ResponseWriter, r *http.Request) {

}

