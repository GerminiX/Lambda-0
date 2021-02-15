package routers

import (
	"github.com/gorilla/mux"
	base_common "github.com/lambda-0/base-common/base-common"
	"github.com/lambda-0/base-common/controllers"
	"github.com/urfave/negroni"
)

func SetupTaskRoutes(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	taskRouter.HandleFunc("/api/v1/tasks/new", controllers.CreateTaskHandler).Methods("POST")
	taskRouter.HandleFunc("/api/v1/tasks/{id}", controllers.UpdateTaskHandler).Methods("PUT")
	taskRouter.HandleFunc("/api/v1/tasks", controllers.GetTasksHandler).Methods("GET")
	taskRouter.HandleFunc("/api/v1/tasks/{id}", controllers.GetTaskByIdHandler).Methods("GET")
	taskRouter.HandleFunc("/api/v1/tasks/users/{id}", controllers.GetTaskByUserHandler).Methods("GET")
	taskRouter.HandleFunc("/api/v1/tasks/{id}", controllers.DeleteTaskHandler).Methods("DELETE")
	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(base_common.Authorize),
		negroni.Wrap(taskRouter),
	))
	return router
}