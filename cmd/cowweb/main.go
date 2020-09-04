package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

const (
	defaultMoosage = "Moo!"
	envCowwebWaitMillisecondKey = "COWWEB_WAIT_MILLISECOND"
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

func waitWithLoad(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			return
		default:
		}
	}
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
		fmt.Fprintln(w, say)
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
		fmt.Fprintln(w, say)
	})

	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		moo := r.URL.Query().Get("m")
		if moo == "" {
			moo = defaultMoosage
		}
		say, err := cowsay.Say(cowsay.Phrase(moo), randomCowType(),	cowsay.Thinking())
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(w, say)

		wenv := os.Getenv(envCowwebWaitMillisecondKey)
		wms, err := strconv.ParseInt(wenv, 10, 0)
		if err != nil {
			wms = 5
		}
		ctx, _ := context.WithTimeout(context.Background(), time.Duration(wms) * time.Millisecond)
		waitWithLoad(ctx)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

