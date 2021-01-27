package routers

import (
	"github.com/gorilla/mux"
)

func SetupUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.userRegister).Methods("POST")
	router.HandleFunc("/users/login", controllers.userLogin).Methods("POST")
	return router
}
