package routers

import (
	"github.com/gorilla/mux"
	"github.com/lambda-0/base-common/controllers"
)

func SetupUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.UserRegister).Methods("POST")
	router.HandleFunc("/users/login", controllers.UserLogin).Methods("POST")
	return router
}
