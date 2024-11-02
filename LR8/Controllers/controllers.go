package Controllers

import (
	"LR8/CRUD"
	"encoding/json"
	"fmt"
	"net/http"
)

type user_str struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SetupHandlers() {
	http.Handle("/hello", Log(http.HandlerFunc(helloHandler)))
	http.Handle("/", Log(http.HandlerFunc(dataHandler)))

	//CRUD
	http.Handle("/users", Log(http.HandlerFunc(usersHandler)))
	http.Handle("/users/1", Log(http.HandlerFunc(getUserHandler)))
	http.Handle("/create", Log(http.HandlerFunc(createUserHandler)))
	http.Handle("/delete/1", Log(http.HandlerFunc(deleteUser)))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
}
func dataHandler(w http.ResponseWriter, r *http.Request) {
	data := decodeData(w, r)

	fmt.Println("ID: ", data.ID)
	fmt.Println("Name: ", data.Name)

	fmt.Fprint(w, "успешно")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	CRUD.GetUsers()
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := decodeData(w, r).ID
	CRUD.GetUserById(id)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	CRUD.CreateUser()
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := decodeData(w, r).ID
	CRUD.DeleteUser(id)
}

func decodeData(w http.ResponseWriter, r *http.Request) user_str {
	decoder := json.NewDecoder(r.Body)
	var user user_str
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return user
}
