package routers

import (
	"github.com/gorilla/mux"
	"github.com/lambda-0/base-common/controllers"
)

func SetupUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.UserRegisterHandler).Methods("POST")
	router.HandleFunc("/users/login", controllers.UserLoginHandler).Methods("POST")
	return router
}
