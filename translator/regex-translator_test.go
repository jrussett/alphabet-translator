package translator_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"alphabet-translator/translator"
)

var _ = Describe("Regex Translator", func() {
	Context("Defaults", func() {
		var rt translator.RegexTranslator

		BeforeEach(func() {
			rt = translator.NewRegexTranslator()
		})

		It("changes hawaiian short vowels", func() {
			Expect(rt.Translate("a")).To(Equal("α"))
			Expect(rt.Translate("e")).To(Equal("ε"))
			Expect(rt.Translate("i")).To(Equal("ι"))
			Expect(rt.Translate("o")).To(Equal("ο"))
			Expect(rt.Translate("u")).To(Equal("υ"))
		})

		It("changes hawaiian long vowels", func() {
			Expect(rt.Translate("ā")).To(Equal(`ᾱ`))
			Expect(rt.Translate("ē")).To(Equal("ε̄"))
			Expect(rt.Translate("ī")).To(Equal("ῑ"))
			Expect(rt.Translate("ō")).To(Equal("ō"))
			Expect(rt.Translate("ū")).To(Equal("ῡ"))
		})

		It("changes hawaiian consonants", func() {
			Expect(rt.Translate("h")).To(Equal("χ"))
			Expect(rt.Translate("k")).To(Equal("κ"))
			Expect(rt.Translate("l")).To(Equal("λ"))
			Expect(rt.Translate("m")).To(Equal("μ"))
			Expect(rt.Translate("n")).To(Equal("ν"))
			Expect(rt.Translate("p")).To(Equal("π"))
			Expect(rt.Translate("w")).To(Equal("ω"))
			Expect(rt.Translate("ʻ")).To(Equal("ξ"))
		})

		It("changes the whole input, not just one letter", func() {
			source := "hānau kūʻokoʻa ʻia nā kānaka apau loa, a ua kau like ka hanohano a me nā pono kīwila ma luna o kākou pākahi."
			result := "χᾱναυ κῡξοκοξα ξια νᾱ κᾱνακα απαυ λοα, α υα καυ λικε κα χανοχανο α με νᾱ πονο κῑωιλα μα λυνα ο κᾱκου πᾱκαχι."
			Expect(rt.Translate(source)).To(Equal(result))
		})
	})
})
