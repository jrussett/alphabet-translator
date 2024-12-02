package translator

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Translator interface {
	Translate(string) string
}

type PotatoTranslator struct{}

func (pt PotatoTranslator) Translate(input string) (output string) {
	output = "potato"
	return output
}
