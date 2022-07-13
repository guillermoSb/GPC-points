package gl

import (
	"encoding/binary"
	"log"
	"os"
)


type Renderer struct {
	width, height uint32
	clearColor Color
	color Color
	pixels [][]Color
	vpX,vpY,vpWidth,vpHeight uint32

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
	r.GlViewPort(0,0,r.width,r.height)
}

func (r * Renderer) GlClearColor(red, green, blue float32) {
	color, err := ColorFromRGB(red, green, blue)
	if err != nil {
		log.Fatal(err)
	}
	r.clearColor = color
}

func (r * Renderer) GlColor(red, green, blue float32) {
	color, err := ColorFromRGB(red, green, blue)
	if err != nil {
		log.Fatal(err)
	}
	r.color = color
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
	// TODO: Implement new way of drawing points
	// Get the row for the point
	if point.X >= float32(r.width) || point.X < 0{
		return;
	}
	if point.Y >= float32(r.height) || point.Y < 0{
		return;
	}

	// row := int(math.Floor(float64(point.X) * float64(r.width - 1)))
	// column := int(math.Floor(float64(point.Y) * float64(r.height - 1)))
	r.pixels[int(point.Y)][int(point.X)] = r.color
}

func (r *Renderer) GlViewPort(posX, posY, width, height uint32) {
	r.vpX = posX
	r.vpY = posY
	r.vpWidth = width
	r.vpHeight = height
}

func (r *Renderer) GlClearViewport(color Color) {
	for x := int(r.vpX); x < (int(r.vpX) + int(r.vpWidth)); x++ {
		for y := int(r.vpY); y < (int(r.vpY) + int(r.vpHeight)); y++ {
			r.pixels[y][x] = color
		}
	}
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
	f.Write(dword(r.width * r.height * 3))	// File Size
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