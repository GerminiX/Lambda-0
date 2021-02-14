package routers

import (
	"github.com/gorilla/mux"
	base_common "github.com/lambda-0/base-common/base-common"
	"github.com/urfave/negroni"
)

func SetupNoteRoutes(router *mux.Router) *mux.Router {
	noteRouter := mux.NewRouter()
	noteRouter.HandleFunc("/api/v1/notes/new", controllers.createNoteHandler).Methods("POST")
	noteRouter.HandleFunc("/api/v1/notes/{id}", controllers.updateNoteHandler).Methods("PUT")
	noteRouter.HandleFunc("/api/v1/notes", controllers.getNotesHandler).Methods("GET")
	noteRouter.HandleFunc("/api/v1/notes/{id}", controllers.getNoteByIdHandler).Methods("GET")
	noteRouter.HandleFunc("/api/v1/notes/tasks/{id}", controllers.getNoteByTaskHandler).Methods("GET")
	noteRouter.HandleFunc("/api/v1/notes/{id}", controllers.deleteNoteHandler).Methods("DELETE")
	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(base_common.Authorize),
		negroni.Wrap(noteRouter),
	))
	return router
}
