package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

var pointsCache = map[[16]byte]Point{}

func main() {
	rec := NewRectangle(6, 4)

	//isto não vai funcionar, pois necessitamos de uma interface RasterImage, o qual deve ter acessõ à []Point
	// DrawPoints(rec)

	adapter := &LineToPointsAdapter{VectorImage: *rec}
	DrawPoints(adapter)
}

type Line struct {
	X1, Y1, X2, Y2 int
}

//Interface fornecida
type VectorImage struct {
	Lines []Line
}

//Interface requerida pelo sistema
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

type LineToPointsAdapter struct {
	VectorImage VectorImage
}

func (ref *LineToPointsAdapter) GetPoints() []Point {
	points := []Point{}

	for _, line := range ref.VectorImage.Lines {
		newPoints := []Point{{X: line.X1, Y: line.Y1}, {X: line.X2, Y: line.Y2}}

		for _, point := range newPoints {
			bytes, _ := json.Marshal(&point)
			hash := md5.Sum(bytes)
			if _, ok := pointsCache[hash]; !ok {
				pointsCache[hash] = point
				points = append(points, point)
			}
		}
	}

	return points
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{
		Lines: []Line{
			Line{0, 0, width, 0},
			Line{0, 0, 0, height},
			Line{0, height, width, height},
			Line{width, 0, width, height},
		},
	}
}

func DrawPoints(owner RasterImage) {
	points := owner.GetPoints()
	for _, point := range points {
		fmt.Println(point)
	}
}
