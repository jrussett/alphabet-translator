package translator

import (
	"regexp"
)

type RegexTranslator struct {
	Glyphs map[string]string
}

func NewRegexTranslator() RegexTranslator {
	rt := RegexTranslator{
		Glyphs: map[string]string{
			"a": "α",
			"e": "ε",
			"i": "ι",
			"o": "ο",
			"u": "υ",

			"ā": "ᾱ",
			"ē": "ε̄",
			"ī": "ῑ",
			"ō": "ō",
			"ū": "ῡ",

			"h": "χ",
			"k": "κ",
			"l": "λ",
			"m": "μ",
			"n": "ν",
			"p": "π",
			"w": "ω",
			"ʻ": "ξ",
		},
	}
	return rt
}

func (rt RegexTranslator) Translate(input string) (output string) {
	output = input

	for src, transformation := range rt.Glyphs {
		re := regexp.MustCompile(src)

		output = re.ReplaceAllLiteralString(output, transformation)
	}

	return output
}
