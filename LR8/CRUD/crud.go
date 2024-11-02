//Реализовать функции для:
// /GET/users ; /GET/users/{id} ; /POST/users ; PUT /users/{id} ; DELETE /users/{id}

package CRUD

import "fmt"

func GetUsers() {
	fmt.Print("users")
}

func GetUserById(id int) {
	fmt.Print("user by id")
	fmt.Print(id)
}

func CreateUser() {
	fmt.Print("create user")
}

func DeleteUser(id int) {
	fmt.Print("delete user")
}
