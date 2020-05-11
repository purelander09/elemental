package main

import (
	"fmt"
	"net/http"

	"github.com/purelander09/elemental/routing"
)

func main() {
	fmt.Println("Starting up Elemental")

	router := routing.CreateRouter()

	http.ListenAndServe(":3333", router)
}
