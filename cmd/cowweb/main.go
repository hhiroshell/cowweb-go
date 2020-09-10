package main

import (
	"log"

	"github.com/hhiroshell/cowweb/pkg/api"
	"github.com/hhiroshell/cowweb/pkg/infrastructure"
)

func main() {
	server := api.NewAPIServer(infrastructure.NewCowsay())
	log.Fatal(server.ListenAndServe())
}

