package config_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"alphabet-translator/config"
)

var _ = Describe("Config", func() {
	var validConfigBytes []byte

	BeforeEach(func() { // TODO: consider doing this in a BeforeSuite
		validConfigBytes = []byte(`
---
name: potato
glyphs:
  a: α
  e: ε
  i: ι
`)
	})

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
				testConfig = validConfigBytes
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

	Describe("NewLanguageConfigFromFile", func() {
		var filePath string

		Context("when the file doesn't exist", func() {
			BeforeEach(func() {
				filePath = "/totally/real/file/path/config.yml"
			})

			It("errors", func() {
				cfg, err := config.NewLanguageConfigFromFile(filePath)
				Expect(err).To(HaveOccurred())
				Expect(cfg).To(BeNil())
			})

			Context("when the file exists and is valid", func() {
				BeforeEach(func() {
					tempFile, err := os.CreateTemp("", "langconfig")
					Expect(err).NotTo(HaveOccurred())

					_, err = tempFile.Write(validConfigBytes)
					Expect(err).NotTo(HaveOccurred())

					filePath = tempFile.Name()

					err = tempFile.Close()
					Expect(err).NotTo(HaveOccurred())

				})

				It("returns a valid LanguageConfig", func() {
					cfg, err := config.NewLanguageConfigFromFile(filePath)
					Expect(err).NotTo(HaveOccurred())

					Expect(cfg.Name).To(Equal("potato"))
					Expect(cfg.Glyphs).To(HaveKeyWithValue("a", "α"))
					Expect(cfg.Glyphs).To(HaveKeyWithValue("e", "ε"))
					Expect(cfg.Glyphs).To(HaveKeyWithValue("i", "ι"))
				})
			})
		})
	})
})
