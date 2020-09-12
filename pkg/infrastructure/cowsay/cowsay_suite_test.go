package cowsay

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPerson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cowsay Suite")
}

type sayOrThink string

const (
	say   sayOrThink = "SAY"
	think sayOrThink = "THINK"
)

func bubble(moosage string, so sayOrThink) (bubble string) {
	bubble += " "
	for i := 0; i < len(moosage)+2; i++ {
		bubble += "_"
	}
	bubble += "\n"
	if so == say {
		bubble += "<"
	} else if so == think {
		bubble += "("
	}
	bubble += " " + moosage + " "
	if so == say {
		bubble += ">"
	} else if so == think {
		bubble += ")"
	}
	bubble += "\n"
	bubble += " "
	for i := 0; i < len(moosage)+2; i++ {
		bubble += "-"
	}
	return
}
