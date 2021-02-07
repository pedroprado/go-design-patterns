package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "This is a BRAVE new world"

	//Sem Flyweight
	fmt.Println("SEM FLYWEIGHT")
	formatted := NewFormattedText(text)
	formatted.Capitalize(10, 15)
	str := formatted.String()
	fmt.Println(str)

	//Com Flyweight
	fmt.Println("COM FLYWEIGHT")
	flyweight := &TextRange{Start: 10, End: 20, Capitalize: true, Italic: false, Bold: false}
	betterFormatted := NewBetterFormattedText(text, []*TextRange{flyweight})

	fmt.Println(betterFormatted.String())
}

//Exemplo: queremos guardar o texto e a formatação fica a cargo dos objetos que consomem o texto
//Dessa forma, temos um FormattedText, que diz qual letra deve ser maiúscula ou minúscula
//Assim, todo texto pode ser guardado sem se preocupar com a formatação
type FormattedText struct {
	plainText  string
	capitalize []bool
}

//Este exemplo é ineficiente, pois, para grandes textos precisariamos uma grande quantidade de memória
//para guardar o slice de boolean

func NewFormattedText(text string) *FormattedText {
	return &FormattedText{plainText: text, capitalize: make([]bool, len(text))}
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		letter := f.plainText[i]
		capitalized := f.capitalize[i]
		if capitalized {
			sb.WriteRune(unicode.ToUpper(rune(letter)))
		} else {
			sb.WriteRune(unicode.ToLower(rune(letter)))
		}
	}
	return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

//----------Implementação com FLYWEIGHT

//Objeto que permite a manipulação da formatação, de forma compacta
//Antes, precisavamos de uma grande lista de boolean
//Agora, apenas um objeto que diz quais as formatações devem ser aplicadas
type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewBetterFormattedText(text string, formatting []*TextRange) *BetterFormattedText {
	return &BetterFormattedText{plainText: text, formatting: formatting}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	textRange := &TextRange{start, end, false, false, false}
	b.formatting = append(b.formatting, textRange)
	return textRange
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(b.plainText); i++ {
		letter := b.plainText[i]

		for _, formatting := range b.formatting {
			//aplica a formatação
			if formatting.Covers(i) {
				if formatting.Capitalize {
					letter = uint8(unicode.ToUpper(rune(letter)))
				} else {
					letter = uint8(unicode.ToUpper(rune(letter)))

				}
			}
		}
		sb.WriteRune(rune(letter))
	}

	return sb.String()
}
