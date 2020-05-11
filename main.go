package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting up Elemental")

	router := routing.createRouter()

	http.ListenAndServe(":3333", router)
}
