package main

import (
	"LR8/Controllers"
	"net/http"
)

func main() {
	Controllers.SetupHandlers()

	http.ListenAndServe(":8090", nil)

}
