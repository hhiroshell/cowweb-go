package main

import (
	"fmt"
	"log"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		say, err := cowsay.Say(
			cowsay.Phrase("Hello"),
			cowsay.Type("default"),
			cowsay.BallonWidth(40),
		)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, say)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
