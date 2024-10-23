package routes

import (
	"TaskHandler/pkg/controllers"

	"github.com/gorilla/mux"
)

func ConfigureRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", controllers.FetchBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.FetchBookByID).Methods("GET")
	router.HandleFunc("/books", controllers.AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.ModifyBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.RemoveBook).Methods("DELETE")

	return router
}
