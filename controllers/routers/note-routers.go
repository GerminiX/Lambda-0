package routers

import (
	"github.com/gorilla/mux"
)

func SetupNoteRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/v1/notes/new", controllers.createNoteHandler).Methods("POST")
	router.HandleFunc("/api/v1/notes/{id}", controllers.updateNoteHandler).Methods("PUT")
	router.HandleFunc("/api/v1/notes", controllers.getNotesHandler).Methods("GET")
	router.HandleFunc("/api/v1/notes/{id}", controllers.getNoteByIdHandler).Methods("GET")
	router.HandleFunc("/api/v1/notes/tasks/{id}", controllers.getNoteByTaskHandler).Methods("GET")
	router.HandleFunc("/api/v1/notes/{id}", controllers.deleteNoteHandler).Methods("DELETE")
	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(base_common.userAuthorizeHandler),
		negroni.Wrap(noteRouter),
	))
	return router
}
