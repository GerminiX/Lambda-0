package routers

import (
	"github.com/gorilla/mux"
	base_common "github.com/lambda-0/base-common/base-common"
	"github.com/urfave/negroni"
)

func SetupTaskRoutes(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	taskRouter.HandleFunc("/api/v1/tasks/new", controllers.createTaskHandler).Methods("POST")
	taskRouter.HandleFunc("/api/v1/tasks/{id}", controllers.updateTaskHandler).Methods("PUT")
	taskRouter.HandleFunc("/api/v1/tasks", controllers.getTasksHandler).Methods("GET")
	taskRouter.HandleFunc("/api/v1/tasks/{id}", controllers.getTaskByIdHandler).Methods("GET")
	taskRouter.HandleFunc("/api/v1/tasks/users/{id}", controllers.getTaskByUserHandler).Methods("GET")
	taskRouter.HandleFunc("/api/v1/tasks/{id}", controllers.deleteTaskHandler).Methods("DELETE")
	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(base_common.Authorize),
		negroni.Wrap(taskRouter),
	))
	return router
}