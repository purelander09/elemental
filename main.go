package main

import (
	"fmt"
	"net/http"

	"github.com/purelander09/elemental/config"
	"github.com/purelander09/elemental/routing"
)

func main() {
	fmt.Println("Starting up Elemental")

	var config = config.GetConfiguration("config.yml")

	router := routing.CreateRouter()

	err := http.ListenAndServe(config.Server.Port, router)

	fmt.Println(err)
}
