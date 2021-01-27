package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router  {
	router := mux.NewRouter().StrictSlash(false)
	router = SetupUserRoutes(router)
	router = SetupTaskRoutes(router)
	router = SetupNoteRoutes(router)
	return router
}
