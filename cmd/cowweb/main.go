package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

const (
	defaultMoosage = "Moo!"
	//defaultWaitMillisecond = 5
	//envCowwebWaitMillisecondKey = "COWWEB_WAIT_MILLISECOND"
	defaultLoad = 1000
	envCowwebLoadKey = "COWWEB_LOAD"
)

var cows []string
var load int

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

	lenv := os.Getenv(envCowwebLoadKey)
	l, err := strconv.ParseInt(lenv, 10, 0)
	if err != nil {
		l = defaultLoad
	}
	load = int(l)
}

func randomCowType() cowsay.Option {
	return cowsay.Type(cows[rand.Intn(len(cows))])
}

func randomCowTypeWithLoad() cowsay.Option {
	//randInts := []int{}
	for i := 0; i < load; i++ {
		for j := 0; j < load; j++ {
			//randInts = append(randInts, rand.Intn(len(cows)))
			rand.Intn(len(cows))
		}
	}
	//return cowsay.Type(cows[randInts[rand.Intn(load)]])
	return cowsay.Type(cows[rand.Intn(len(cows))])
}

//func waitWithLoad(ctx context.Context) {
//	for {
//		select {
//		case <- ctx.Done():
//			return
//		default:
//		}
//	}
//}

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
		say, err := cowsay.Say(cowsay.Phrase(moo), randomCowTypeWithLoad())
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(w, say)

		//wenv := os.Getenv(envCowwebWaitMillisecondKey)
		//wms, err := strconv.ParseInt(wenv, 10, 0)
		//if err != nil {
		//	wms = defaultWaitMillisecond
		//}
		//ctx, _ := context.WithTimeout(context.Background(), time.Duration(wms) * time.Millisecond)
		//waitWithLoad(ctx)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

