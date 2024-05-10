package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"pizzaStoreAPI/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func initHeadders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetAllPizzas(w http.ResponseWriter, r *http.Request) {
	initHeadders(w)
	log.Println("GET info about all pizzas")
	w.WriteHeader(200)                   //  Установка и отправка статус кода для запроса
	json.NewEncoder(w).Encode(models.Db) //  Маршалинг данных(.Encode(DB)) и отправка клиенту(NewEncoder(w)) в виде json
}

func GetPizzaById(w http.ResponseWriter, r *http.Request) {
	initHeadders(w)
	id, err := strconv.Atoi(mux.Vars(r)["id"]) //  Считать параметры запроса
	if err != nil {
		log.Println("GET info about pizza with invalid id: ", err)
		msg := models.Message{Message: "this ID can not casting to int"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	log.Println("GET info about pizza with id:", id)
	pizza, ok := models.FindPizzaByID(id)
	if !ok {
		w.WriteHeader(404)
		msg := models.Message{Message: "pizza with id does not exist, try another id"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(pizza)
}

func CreatePizza(w http.ResponseWriter, r *http.Request) {
	initHeadders(w)
	log.Println("Add new pizza in menu ...")
	var pizza models.Pizza
	err := json.NewDecoder(r.Body).Decode(&pizza) //  Десереализуем json в модель пиццы
	if err != nil {
		msg := models.Message{Message: "Invalid data"} //  Формируем сообщение об ошибке данных
		w.WriteHeader(400)                             //  Установка статус кода
		json.NewEncoder(w).Encode(msg)                 //  Отправка ответа
		return
	}
	newPizzaId := len(models.Db) + 1 // create id
	pizza.ID = newPizzaId
	models.Db = append(models.Db, pizza)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(pizza)
	log.Printf("new pizza added with id: %v, title: %v", pizza.ID, pizza.Title)
}

func UpdatePizzaById(w http.ResponseWriter, r *http.Request) {
	initHeadders(w)
	log.Println("Updating pizza ...")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("GET info about pizza with invalid id: ", err)
		msg := models.Message{Message: "this ID can not casting to int"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	pizza, ok := models.FindPizzaByID(id)
	if !ok {
		log.Println("pizza not found with id:", id)
		w.WriteHeader(404)
		msg := models.Message{Message: "pizza with id does not exist, try another id"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&pizza)
	if err != nil {
		msg := models.Message{Message: "Invalid data"} //  Формируем сообщение об ошибке данных
		w.WriteHeader(400)                             //  Установка статус кода
		json.NewEncoder(w).Encode(msg)                 //  Отправка ответа
		return
	}
	// TODO: Update pizza in DB
	json.NewEncoder(w).Encode(pizza)
	log.Printf("pizza was updated with id: %v, title: %v", pizza.ID, pizza.Title)

}

func DeletePizzaById(w http.ResponseWriter, r *http.Request) {
	initHeadders(w)
	log.Println("Deleting pizza ...")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("GET info about pizza with invalid id: ", err)
		msg := models.Message{Message: "this ID can not casting to int"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	pizza, ok := models.FindPizzaByID(id)
	if !ok {
		log.Println("pizza not found with id:", id)
		w.WriteHeader(404)
		msg := models.Message{Message: "pizza with id does not exist, try another id"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	// TODO: Delete pizza out of DB
	w.WriteHeader(200)
	msg := models.Message{Message: "successfuly deleted"}
	json.NewEncoder(w).Encode(msg)
	log.Printf("pizza with id %v deleted", pizza.ID)

}
