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
	userRepo := &data.UserRepository{Col: userCol}
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
	var dataResource LoginResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		base_common.DisplayAppError(w,
			err,
			"Invalid login data",
			500,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email: loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	userCol := context.Collection(models.UserCollection)
	userRepo := &data.UserRepository{Col: userCol}
	if user, err := userRepo.Login(loginUser); err != nil {
		base_common.DisplayAppError(w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	}else {
		token, err := base_common.GenerateToken(user.Email, "member")
		if err != nil {
			base_common.DisplayAppError(w,
				err,
				"Error while generating the access token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User: user,
			Token: token,
		}
		resp, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			base_common.DisplayAppError(w,
				err,
				"An unexpected error has occured",
				500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}

}

