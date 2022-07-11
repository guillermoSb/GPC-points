package main // Main package
import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
)

// Encoding binary
//

func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	renderer := Renderer{width: 40, height: 40}
	renderer.glInit()
	renderer.glColor(0,1,0)
	renderer.glPoint(Point{0.75,0.75})
	renderer.glColor(0,0,1)
	renderer.glPoint(Point{1,0.99})
	renderer.glFinish("out.bmp")
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

type Renderer struct {
	width, height uint32
	clearColor Color
	color Color
	pixels [][]Color
	
}


// Initialize the renderer
func (r * Renderer) glInit() {
		r.clearColor = Color{0.4,0,0}	// Default color is black
		r.color = Color {1,1,1}	// Default color is white
		r.glClear()	// Clear the image
}

func (r * Renderer) glClearColor(red float32, green float32, blue float32) {
	// TODO: Validate that the color is within the expected ranges
	r.clearColor = Color{red,green,blue}
}

func (r * Renderer) glColor(red float32, green float32, blue float32) {
	// TODO: Validate that the color is within the expected ranges
	r.color = Color{red,green,blue}
}


func (r * Renderer) glClear() {
	for i := 0; i < int(r.width); i++ {
		row := []Color{}
		for j := 0; j < int(r.height); j++ {
			row= append(row, r.clearColor)
		}
		r.pixels = append(r.pixels, row)
	}
}


func (r * Renderer) glPoint(point Point) {
	// Get the row for the point

	row := int(math.Floor(float64(point.x) * float64(r.width - 1)))
	column := int(math.Floor(float64(point.y) * float64(r.height - 1)))
	r.pixels[row][column] = r.color

}

func (r * Renderer) glFinish(fileName string) {
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




	// fmt.Println(binary.BigEndian.Uint16([]byte("BM")))
}

type Color struct {
	r, g, b float32
}

func (c * Color ) bytes()[]byte {
	red := uint8(c.r * 255)
	green := uint8(c.g * 255)
	blue := uint8(c.b * 255)
	return []byte{blue, green, red}
}



type Point struct {
	x float32
	y float32
}