package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hhiroshell/cowweb/pkg/domain/service"
)

func NewAPIServer(cowsay service.CowService, port int) (*http.Server, error) {
	if cowsay == nil {
		return nil, fmt.Errorf("illegal argument. cowsay == nil")
	}
	if port <= 0 {
		return nil, fmt.Errorf("illegal argument. port <= 0")
	}
	h := &handlers{cowsay: cowsay}
	http.HandleFunc("/", h.say)
	http.HandleFunc("/say", h.say)
	http.HandleFunc("/think", h.think)
	http.HandleFunc("/health", h.health)
	return &http.Server{Addr: ":" + strconv.Itoa(port)}, nil
}

type handlers struct {
	cowsay service.CowService
}

func (a *handlers) say(w http.ResponseWriter, r *http.Request) {
	moosage := r.URL.Query().Get("moosage")
	cow, err := a.cowsay.Say(moosage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Fprintln(w, cow)
}

func (a *handlers) think(w http.ResponseWriter, r *http.Request) {
	moosage := r.URL.Query().Get("moosage")
	cow, err := a.cowsay.Think(moosage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Fprintln(w, cow)
}

func (a *handlers) health(w http.ResponseWriter, r *http.Request) {
	return
}
