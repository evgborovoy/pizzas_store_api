package main

import (
	"log"
	"net/http"
	"os"
	"pizzaStoreAPI/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	port             string
	apiPrefix        string = "/api/v1"
	manyPizzasPrefix string = apiPrefix + "/pizzas"
	pizzaPrefix      string = apiPrefix + "/pizza"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file: ", err)
	}

	port = os.Getenv("app_port")
}

func main() {
	log.Println("Trying to start REST API pizza on port:", port)
	router := mux.NewRouter()
	utils.BuildManyPizzasResource(router, manyPizzasPrefix)
	utils.BuildPizzaResource(router, pizzaPrefix)
	log.Println("Router configured successfuly. Let's Go!")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
