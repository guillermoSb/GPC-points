package gl

import (
	"encoding/binary"
	"log"
	"os"
)

// Structure that defines the properties of a Renderer
// Properties:
// - width, height: Dimensions of the image
// - clearColor: Color to use as a clear
// - color: Color to use as a color for pixels
// - pixels: Slice of pixels to draw
// - vpx,vpY,vpWidth,vpHeight: Viewport offset and dimensions
type Renderer struct {
	width, height uint32
	clearColor Color
	color Color
	pixels [][]Color
	vpX,vpY,vpWidth,vpHeight uint32
}

// Initialize the renderer
// Parameters:
// - width: Width of the renderer
// - height: Height of the renderer
func (r * Renderer) GlInit(width, height uint32) {
	r.glCreateWindow(width, height)
}
// Creates a new window 
// Parameters:
// - width: Width of the window
// - height: Height of the window
func (r * Renderer) glCreateWindow(width, height uint32) {
	r.width = width
	r.height = height
	r.pixels = [][]Color{}
	r.GlViewPort(0,0,r.width,r.height)
}

// Sets the clearColor for the renderer
// Parameters:
// - red: red value for the clear color.
// - green: green value for the clear color.
// - blue: blue value for the clear color.
func (r * Renderer) GlClearColor(red, green, blue float32) {
	color, err := ColorFromRGB(red, green, blue)
	if err != nil {
		log.Fatal(err)
	}
	r.clearColor = color
}

// Sets the color for the renderer
// Parameters:
// - red: red value for the color.
// - green: green value for the color.
// - blue: blue value for the color.
func (r * Renderer) GlColor(red, green, blue float32) {
	color, err := ColorFromRGB(red, green, blue)
	if err != nil {
		log.Fatal(err)
	}
	r.color = color
}

// Clears the renderer with the clearColor
func (r * Renderer) GlClear() {
	for i := 0; i < int(r.width); i++ {
		row := []Color{}
		for j := 0; j < int(r.height); j++ {
			row= append(row, r.clearColor)
		}
		r.pixels = append(r.pixels, row)
	}
}

// Draws a point on the renderer
// Parameters:
// - point: the point to draw
func (r * Renderer) GlPoint(point Point) {
	// Get the row for the point
	if (point.X >= float32(r.width) || point.X < 0) || (point.Y >= float32(r.height) || point.Y < 0){
		return;
	}
	r.pixels[int(point.Y)][int(point.X)] = r.color
}

// Draws a point on the viewport
// Parameters:
// - point: the point to draw
func (r * Renderer) GlViewPortPoint(point Point) {
	if point.X < -1 || point.X > 1 || point.Y < -1 || point.Y > 1 {
		return;
	}
	x := (point.X + 1) * ((float32(r.vpWidth - 1) / 2 )) + float32(r.vpX)
	y := (point.Y + 1) * ((float32(r.vpHeight - 1) / 2 )) + float32(r.vpY)
	r.GlPoint(Point{x,y})
}

// Sets the properties for a viewport
// Parameters:
// - posX: position x for the viewport
// - posY: position y for the viewport
// - width: width for the viewport
// - height: height for the viewport
func (r *Renderer) GlViewPort(posX, posY, width, height uint32) {
	r.vpX = posX
	r.vpY = posY
	r.vpWidth = width
	r.vpHeight = height
}

// Clears the viewport with a given color
// Parameters:
// - color: color value for the clear color.
func (r *Renderer) GlClearViewport(color Color) {
	for x := int(r.vpX); x < (int(r.vpX) + int(r.vpWidth)); x++ {
		for y := int(r.vpY); y < (int(r.vpY) + int(r.vpHeight)); y++ {
			r.pixels[y][x] = color
		}
	}
}

// Finishes rendering the image
// Parameters:
// - fileName: Name of the file to output.
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