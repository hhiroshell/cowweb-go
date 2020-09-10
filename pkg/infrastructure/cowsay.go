package infrastructure

import (
	"math/rand"

	cowsay "github.com/Code-Hex/Neo-cowsay"

	"github.com/hhiroshell/cowweb/pkg/domain/service"
)

const (
	defaultMoosage = "Moo!"
	defaultLoad = 1024
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
}

type Cowsay struct {}

func NewCowsay() service.CowService {
	return &Cowsay{}
}

func (c *Cowsay) Say(moosage string) (string, error) {
	if moosage == "" {
		moosage = defaultMoosage
	}
	return cowsay.Say(cowsay.Phrase(moosage), randomCowType())
}

func (c *Cowsay) Think(moosage string) (string, error) {
	if moosage == "" {
		moosage = defaultMoosage
	}
	return cowsay.Say(cowsay.Phrase(moosage), randomCowType(), cowsay.Thinking())
}

type SlowCowsay struct {
	load int
}

func NewSlowCowsay(load int) service.CowService {
	if load <= 0 {
		load = defaultLoad
	}
	return &SlowCowsay{load: load}
}

func (c *SlowCowsay) Say(moosage string) (string, error) {
	if moosage == "" {
		moosage = defaultMoosage
	}
	for i := 0; i < load; i++ {
		for j := 0; j < load; j++ {
			rand.Intn(len(cows))
		}
	}
	return cowsay.Say(cowsay.Phrase(moosage), randomCowType())
}

func (c *SlowCowsay) Think(moosage string) (string, error) {
	if moosage == "" {
		moosage = defaultMoosage
	}
	for i := 0; i < load; i++ {
		for j := 0; j < load; j++ {
			rand.Intn(len(cows))
		}
	}
	return cowsay.Say(cowsay.Phrase(moosage), randomCowType(), cowsay.Thinking())
}

func randomCowType() cowsay.Option {
	return cowsay.Type(cows[rand.Intn(len(cows))])
}
