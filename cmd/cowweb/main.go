package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

const (
	defaultMoosage = "Moo!"
)

var cows []string

func init() {
	var infelicities = [3]string{"head-in", "telebears", "sodomized"}
	contains := func(s string) bool {
		for _, inf := range infelicities {
			if s == inf {
				return true
			}
		}
		return false
	}
	for _, c := range cowsay.Cows() {
		if contains(c) {
			continue
		}
		cows = append(cows, c)
	}
}

func randomCowType() cowsay.Option {
	return cowsay.Type(cows[rand.Intn(len(cows))])
}

func main() {
	http.HandleFunc("/say", func(w http.ResponseWriter, r *http.Request) {
		moo := r.URL.Query().Get("m")
		if moo == "" {
			moo = defaultMoosage
		}
		say, err := cowsay.Say(cowsay.Phrase(moo), randomCowType())
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, say)
	})

	http.HandleFunc("/think", func(w http.ResponseWriter, r *http.Request) {
		moo := r.URL.Query().Get("m")
		if moo == "" {
			moo = defaultMoosage
		}
		say, err := cowsay.Say(cowsay.Phrase(moo), randomCowType(),	cowsay.Thinking())
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, say)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

