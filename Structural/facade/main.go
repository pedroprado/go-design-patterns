package main

import "fmt"

func main() {

	c := NewConsole()
	u := c.CharacterAt(1)

	fmt.Println(u)
}

//Buffers e Viewports. Usuário não está interessado em mexer com esses detalhes
//Para uma usabilidade mais simples, construimos uma facade (Console)
type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width int, height int) *Buffer {
	return &Buffer{width, height, make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(index)
}

//Console: um sistema complexo que contém vários buffers e portas
type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

func (c *Console) CharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}
