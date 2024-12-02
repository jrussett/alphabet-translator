package translator_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"alphabet-translator/translator"
)

var _ = Describe("Translator", func() {
	Context("Potato Translator", func() {
		var pt *translator.PotatoTranslator

		BeforeEach(func() {
			pt = &translator.PotatoTranslator{}
		})

		It("Returns a dummy value for any input", func() {
			dummyValue := "potato"

			for _, givenValue := range []string{"rutabaga", "carrot", "broccoli"} {
				Expect(pt.Translate(givenValue)).To(Equal(dummyValue))
			}
		})

	})

})
