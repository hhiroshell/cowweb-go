package cowsay

import (
	"math/rand"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/hhiroshell/cowweb/pkg/domain/service"
)

type SlowCowsay struct {
	load int
}

func NewSlowCowsay(load int) service.CowService {
	return &SlowCowsay{load: load}
}

func (c *SlowCowsay) Say(moosage string) (string, error) {
	for i := 0; i < c.load; i++ {
		for j := 0; j < c.load; j++ {
			rand.Intn(len(cows))
		}
	}
	return cowsay.Say(cowsay.Phrase(moosage), randomCowType())
}

func (c *SlowCowsay) Think(moosage string) (string, error) {
	for i := 0; i < c.load; i++ {
		for j := 0; j < c.load; j++ {
			rand.Intn(len(cows))
		}
	}
	return cowsay.Say(cowsay.Phrase(moosage), randomCowType(), cowsay.Thinking())
}
