package gl

import (
	"encoding/binary"
	"log"
	"math"
	"os"
)


type Renderer struct {
	width, height uint32
	clearColor Color
	color Color
	pixels [][]Color
	viewPort [][]Color

}

// Initialize the renderer
func (r * Renderer) GlInit() {
		r.clearColor = Color{0.4,0,0}	// Default color is black
		r.color = Color {1,1,1}	// Default color is white
}

func (r * Renderer) GlCreateWindow(width, height uint32) {
	r.width = width
	r.height = height
	r.pixels = [][]Color{}
	r.viewPort = [][]Color{}
}

func (r * Renderer) GlClearColor(red float32, green float32, blue float32) {
	// TODO: Validate that the color is within the expected ranges
	r.clearColor = Color{red,green,blue}
}

func (r * Renderer) GlColor(red float32, green float32, blue float32) {
	// TODO: Validate that the color is within the expected ranges
	r.color = Color{red,green,blue}
}

func (r * Renderer) GlViewPort(point Point, width, height float32) {
	// TODO: Validate that the viewport is within the expected ranges
}

func (r * Renderer) GlClear() {
	for i := 0; i < int(r.width); i++ {
		row := []Color{}
		for j := 0; j < int(r.height); j++ {
			row= append(row, r.clearColor)
		}
		r.pixels = append(r.pixels, row)
	}
}

func (r * Renderer) GlPoint(point Point) {
	// Get the row for the point

	row := int(math.Floor(float64(point.X) * float64(r.width - 1)))
	column := int(math.Floor(float64(point.Y) * float64(r.height - 1)))
	r.pixels[row][column] = r.color

}

func (r * Renderer) GlFinish(fileName string) {
	// Attempt to open the file
	f, err := os.Create(fileName)

	// Check if the file was successfully created
	if err != nil {
		log.Fatal(err)
	}
	// Example, writing 5 x 5 image
	defer f.Close()	// Close the file when the process is done
	f.Write([]byte("B"))
	f.Write([]byte("M"))
	f.Write([]byte{0, 0, 0, 0})	// File Size
	f.Write([]byte{0, 0})	// Reserved
	f.Write([]byte{0, 0})	// Reserved
	f.Write([]byte{54, 0, 0, 0 })	// ?
	f.Write([]byte{40, 0, 0, 0})	// Header Size
	f.Write(dword(r.width))		// Width
	f.Write(dword(r.height))		// Height
	f.Write([]byte{1, 0})	// Plane
	f.Write([]byte{24, 0})	// BPP
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	// Pixel Data
	for i := 0; i < int(r.width); i++ {
		for j := 0; j < int(r.height); j++ {
			f.Write(r.pixels[i][j].bytes())
		}
	}
}

// ****************************************************************
// Utils and Structures
// ****************************************************************

type Color struct {
	R, G, B float32
}

func (c * Color ) bytes()[]byte {
	red := uint8(c.R * 255)
	green := uint8(c.G * 255)
	blue := uint8(c.B * 255)
	return []byte{blue, green, red}
}

type Point struct {
	X float32
	Y float32
}

func word(value uint16) []byte {
	bs := make([]byte,2)
	binary.LittleEndian.PutUint16(bs,value)
	return bs
}

func dword(value uint32) []byte {
	bs := make([]byte,4)
	binary.LittleEndian.PutUint32(bs,value)
	return bs
}