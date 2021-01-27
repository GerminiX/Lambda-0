package routers

import (
	"github.com/gorilla/mux"
)

func SetupTaskRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/v1/tasks/new", controllers.createTaskHandler).Methods("POST")
	router.HandleFunc("/api/v1/tasks/{id}", controllers.updateTaskHandler).Methods("PUT")
	router.HandleFunc("/api/v1/tasks", controllers.getTasksHandler).Methods("GET")
	router.HandleFunc("/api/v1/tasks/{id}", controllers.getTaskByIdHandler).Methods("GET")
	router.HandleFunc("/api/v1/tasks/users/{id}", controllers.getTaskByUserHandler).Methods("GET")
	router.HandleFunc("/api/v1/tasks/{id}", controllers.deleteTaskHandler).Methods("DELETE")
	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(base_common.userAuthorizeHandler),
		negroni.Wrap(taskRouter),
	))
	return router
}