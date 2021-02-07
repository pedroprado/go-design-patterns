package main

import "fmt"

//Sem Brigde: duas dimensões, da Forma e da Renderização.
type Shape interface {
	RenderShape()
}

type CircleVectorRenderer struct {
	Radius float64
}

func (ref *CircleVectorRenderer) RenderShape() {
	fmt.Println("renderizando vetor")
}

type CircleRasterRenderer struct {
	Radius float64
}

func (ref *CircleRasterRenderer) RenderShape() {
	fmt.Println("renderizando raster")
}

type RectangleVectorRenderer struct {
	width, heigth int
}

func (ref *RectangleVectorRenderer) Render() {
	fmt.Println("renderizando vetor")
}

type RectangleRasterRenderer struct {
	width, heigth int
}

func (ref *RectangleRasterRenderer) Render() {
	fmt.Println("renderizando raster")
}

//Uso do Bridge: apenas uma dimensão (Forma) e a outra é transformada em uma dependência
//Uso de composição
type Renderer interface {
	Render()
}

type RasterRenderer struct{}

func (ref *RasterRenderer) Render() {}

type VectorRenderer struct{}

func (ref *VectorRenderer) Render() {}

//A única dimensão aqui é a da Forma. A dimensão da Renderização é eliminada e usada como dependência na Forma.
type Circle struct {
	Radius   float64
	Renderer Renderer
}

func (c *Circle) RenderShape() {
	c.Renderer.Render()
}

type Rectangle struct {
	Width, Heigth int
	Renderer      Renderer
}

func (r *Rectangle) RenderShape() {
	r.Renderer.Render()
}
