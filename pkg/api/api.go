package api

import (
	"fmt"
	"net/http"

	"github.com/hhiroshell/cowweb/pkg/domain/service"
)

const defaultMoosage = "Moo!"

func NewAPIServer(cowsay service.CowService) *http.Server {
	h := &handlers{cowsay: cowsay}
	http.HandleFunc("/say", h.say)
	http.HandleFunc("/think", h.think)
	return &http.Server{Addr: ":8080"}
}

type handlers struct {
	cowsay service.CowService
}

func (a *handlers) say(w http.ResponseWriter, r *http.Request) {
	moosage := r.URL.Query().Get("m")
	if moosage == "" {
		moosage = defaultMoosage
	}
	cow, err := a.cowsay.Say(moosage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Fprintln(w, cow)
}

func (a *handlers) think(w http.ResponseWriter, r *http.Request) {
	moosage := r.URL.Query().Get("m")
	if moosage == "" {
		moosage = defaultMoosage
	}
	cow, err := a.cowsay.Think(moosage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Fprintln(w, cow)
}

