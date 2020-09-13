package cowsay

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the cowsay package", func() {
	Context("slowCowsay.Say()", func() {
		When("the moosage is blank string", func() {
			It("returns a cow aa with default moosage", func() {
				sc := &SlowCowsay{}
				cow, _ := sc.Say("")
				Expect(strings.Contains(cow, bubble(defaultMoosage, say))).To(BeTrue())
			})
		})
		When("the moosage is not blank string", func() {
			It("returns a cow aa with the specified moosage", func() {
				moosage := "this is a test moosage !!"
				sc := &SlowCowsay{}
				cow, _ := sc.Say(moosage)
				Expect(strings.Contains(cow, bubble(moosage, say))).To(BeTrue())
			})
		})
		Measure("it should do same thing slower than the normal one", func(b Benchmarker) {
			expect := bubble(defaultMoosage, say)
			normal := b.Time("runtime", func() {
				c := &Cowsay{}
				cow, _ := c.Say("")
				Expect(strings.Contains(cow, expect)).To(BeTrue())
			})
			slow := b.Time("runtime", func() {
				sc := &SlowCowsay{load: 512}
				cow, _ := sc.Say("")
				Expect(strings.Contains(cow, expect)).To(BeTrue())
			})
			Ω(slow.Seconds()).Should(BeNumerically(
				">", normal.Seconds()), "slowCowsay.Say() should take longer than the normal one.")
		}, 10)
	})
	Context("slowCowsay.Think()", func() {
		When("the moosage is blank string", func() {
			It("returns a cow aa with default moosage", func() {
				sc := &Cowsay{}
				cow, _ := sc.Think("")
				Expect(strings.Contains(cow, bubble(defaultMoosage, think))).To(BeTrue())
			})
		})
		When("the moosage is not blank string", func() {
			It("returns a cow aa with the specified moosage", func() {
				moosage := "this is a test moosage !!"
				sc := &Cowsay{}
				cow, _ := sc.Think(moosage)
				Expect(strings.Contains(cow, bubble(moosage, think))).To(BeTrue())
			})
		})
		Measure("it should do same thing slower than the normal one", func(b Benchmarker) {
			expect := bubble(defaultMoosage, think)
			normal := b.Time("runtime", func() {
				c := &Cowsay{}
				cow, _ := c.Think("")
				Expect(strings.Contains(cow, expect)).To(BeTrue())
			})
			slow := b.Time("runtime", func() {
				sc := &SlowCowsay{load: 512}
				cow, _ := sc.Think("")
				Expect(strings.Contains(cow, expect)).To(BeTrue())
			})
			Ω(slow.Seconds()).Should(BeNumerically(
				">", normal.Seconds()), "slowCowsay.Think() should take longer than the normal one.")
		}, 10)
	})
})
