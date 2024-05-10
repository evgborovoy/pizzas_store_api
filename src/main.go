package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func init() {
	pizza1 := Pizza{ID: 1, Size: 30, Title: "Margarita", Price: 300}
	pizza2 := Pizza{ID: 2, Size: 26, Title: "Pepperoni", Price: 200}
	pizza3 := Pizza{ID: 3, Size: 35, Title: "BBQ", Price: 450}
	db = append(db, pizza1, pizza2, pizza3)
}

var (
	//  Порт запуска приложения
	port string = "8080"
	//  База данных
	db []Pizza
)

// Модель даннных
type Pizza struct {
	ID    int     `json:"id"`
	Size  int     `json:"size"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

// Функция для поиска пиццы по id
func FindPizzaByID(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range db {
		if p.ID == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}

//  по адресу /pizzas возвращает список всех доступных пицц к заказу
//  по адресу /pizzas/<id> возващает информацию о пицце с конкретным <id>, если она есть,
//  или информацию, что такое пиццы нет

// Функции которые будут выполняться при дергании ручки
func GetAllPizzas(w http.ResponseWriter, r *http.Request) {
	//  Настройка хэдеров
	w.Header().Set("Content-Type", "application/json")
	log.Println("GET info about all pizzas") //  Логирование запроса
	w.WriteHeader(200)                       //  Установка и отправка статус кода для запроса
	json.NewEncoder(w).Encode(db)            //  Маршалинг данных(.Encode(db)) и отправка клиенту(NewEncoder(w)) в виде json

}

func GetPizzaById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//  Считаем параметры из запроса и конвертируем его в int
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("GET info about pizza with invalid id: ", err)
		msg := ErrorMessage{Message: "this ID can not casting to int"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	log.Println("GET info about pizza with id:", id) //  Логирование запроса
	pizza, ok := FindPizzaByID(id)
	if !ok {
		w.WriteHeader(404)
		msg := ErrorMessage{Message: "pizza with id does not exist, try another id"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(pizza)
}

func main() {
	log.Println("Trying to start REST API pizza")
	//  Инициализируем маршрутизатор
	router := mux.NewRouter()

	// по какому адресу и какая функция будет выполняться
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	router.HandleFunc("/pizzas/{id}", GetPizzaById).Methods("GET")

	log.Println("Router configured successfuly. Let's Go!")
	//  Включаем сервер на постоянное прослушивание
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
