package utils

import (
	"pizzaStoreAPI/pkg/handlers"

	"github.com/gorilla/mux"
)

func BuildManyPizzasResource(router *mux.Router, adr string) {
	router.HandleFunc(adr, handlers.GetAllPizzas).Methods("GET")

}

func BuildPizzaResource(router *mux.Router, adr string) {
	router.HandleFunc(adr+"/{id}", handlers.GetPizzaById).Methods("GET")
	router.HandleFunc(adr, handlers.CreatePizza).Methods("POST")
	router.HandleFunc(adr+"/{id}", handlers.UpdatePizzaById).Methods("PUT")
	router.HandleFunc(adr+"/{id}", handlers.DeletePizzaById).Methods("DELETE")
}
