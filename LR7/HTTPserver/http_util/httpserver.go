package httputil

import (
	"HTTPserver/middleware"
	"encoding/json"
	"fmt"
	"net/http"
)

type data_str struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func setutHandlers() {
	http.Handle("/hello", middleware.Log(http.HandlerFunc(helloHandler)))
	http.Handle("/data", middleware.Log(http.HandlerFunc(dataHandler)))
}

func Start() {
	//var wg sync.WaitGroup
	//mux := http.NewServeMux()
	setutHandlers()

	http.ListenAndServe(":8090", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data data_str
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Более корректное обраблтование ошибки
		return
	}

	fmt.Println("ID: ", data.ID)
	fmt.Println("Name: ", data.Name)

	fmt.Fprint(w, "успешно")
}
