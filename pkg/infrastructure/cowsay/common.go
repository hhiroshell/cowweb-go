package cowsay

import (
	"math/rand"

	cowsay "github.com/Code-Hex/Neo-cowsay"
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
