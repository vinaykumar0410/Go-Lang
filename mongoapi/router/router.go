package router

import (
	"mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/developer", controller.CreateDeveloper).Methods("POST")
	router.HandleFunc("/api/developers", controller.GetDevelopers).Methods("GET")
	router.HandleFunc("/api/developer/{id}", controller.UpdateDeveloperAsLead).Methods("PUT")
	router.HandleFunc("/api/developer/{id}", controller.DeleteDeveloper).Methods("DELETE")
	router.HandleFunc("/api/developers", controller.DeleteAllDeveloper).Methods("DELETE")

	return router
}