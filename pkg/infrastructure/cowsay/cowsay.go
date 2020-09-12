package cowsay

import (
	cowsay "github.com/Code-Hex/Neo-cowsay"

	"github.com/hhiroshell/cowweb/pkg/domain/service"
)

type Cowsay struct{}

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
