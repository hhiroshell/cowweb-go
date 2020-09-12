package cowsay

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the cowsay package", func() {
	Context("cowsay.Say()", func() {
		When("the moosage is blank string", func() {
			It("returns a cow aa with default moosage", func() {
				c := &Cowsay{}
				cow, _ := c.Say("")
				Expect(strings.Contains(cow, bubble(defaultMoosage, say))).To(BeTrue())
			})
		})
		When("the moosage is not blank string", func() {
			It("returns a cow aa with the specified moosage", func() {
				moosage := "this is a test moosage !!"
				c := &Cowsay{}
				cow, _ := c.Say(moosage)
				Expect(strings.Contains(cow, bubble(moosage, say))).To(BeTrue())
			})
		})
	})
	Context("cowsay.Think()", func() {
		When("the moosage is blank string", func() {
			It("returns a cow aa with default moosage", func() {
				c := &Cowsay{}
				cow, _ := c.Think("")
				Expect(strings.Contains(cow, bubble(defaultMoosage, think))).To(BeTrue())
			})
		})
		When("the moosage is not blank string", func() {
			It("returns a cow aa with the specified moosage", func() {
				moosage := "this is a test moosage !!"
				c := &Cowsay{}
				cow, _ := c.Think(moosage)
				Expect(strings.Contains(cow, bubble(moosage, think))).To(BeTrue())
			})
		})
	})
})
