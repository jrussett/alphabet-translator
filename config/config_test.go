package config_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"alphabet-translator/config"
)

var _ = Describe("Config", func() {
	Describe("NewLanguageConfigFromBytes", func() {
		var testConfig []byte

		Context("when the bytes are invalid", func() {
			BeforeEach(func() {
				testConfig = []byte(`---potato`)
			})

			It("returns an error", func() {
				cfg, err := config.NewLanguageConfigFromBytes(testConfig)
				Expect(err).To(HaveOccurred())
				Expect(cfg).To(BeNil())
			})
		})

		Context("when the bytes are valid", func() {
			BeforeEach(func() {
				testConfig = []byte(`
---
name: potato
glyphs:
  a: α
  e: ε
  i: ι
`)
			})

			It("returns a valid LanguageConfig", func() {
				cfg, err := config.NewLanguageConfigFromBytes(testConfig)
				Expect(err).NotTo(HaveOccurred())

				Expect(cfg.Name).To(Equal("potato"))
				Expect(cfg.Glyphs).To(HaveKeyWithValue("a", "α"))
				Expect(cfg.Glyphs).To(HaveKeyWithValue("e", "ε"))
				Expect(cfg.Glyphs).To(HaveKeyWithValue("i", "ι"))
			})
		})
	})
})
