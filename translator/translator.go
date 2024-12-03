package translator

type Translator interface {
	Translate(string) string
}

type PotatoTranslator struct{}

func (pt PotatoTranslator) Translate(input string) (output string) {
	output = "potato"
	return output
}
