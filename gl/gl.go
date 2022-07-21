package gl

import (
	"encoding/binary"
	"log"
	"math"
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
func (r * Renderer) GlPoint(point Point, colors ...Color) {
	// Get the row for the point
	if len(colors) > 0 {
		r.GlColor(colors[0].R, colors[0].G, colors[0].B)
	}
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

// Draws a line on the viewport using the bresenham line drawing algorithm/
// Takes two vectors to draw and a list of colors. The first value is going to be used.
func (r *Renderer) GLLine(v0, v1 Point, colors ...Color) {
	x0 := v0.X
	x1 := v1.X
	y0 := v0.Y
	y1 := v1.Y
	
	

	// Dibujar un punto si ambos valores son iguales
	if x0 == x1 && y0 == y1{
		r.GlPoint(v0, colors...)
		return
	}
	
	dy := math.Abs(float64(y1) - float64(y0))
	dx := math.Abs(float64(x1) - float64(x0))

	steep := dy > dx
	// Tiene pendiente mayor a 1 o menor a -1, intercambiar las x por las y, y se dibuja una línea vertical
	if steep {
		tempx0 := x0
		tempx1 := x1
		tempy0 := y0
		tempy1 := y1
		x0 = tempy0
		y0 = tempx0
		x1 = tempy1
		y1 = tempx1
	}
	// Si el punto inicial en X, es mayor que el punto final X, intercambiar los puntos para siempre dibujar de iz a derecha
	if x0 > x1 {
		tempx0 := x0
		tempx1 := x1
		tempy0 := y0
		tempy1 := y1
		x0 = tempx1
		x1 = tempx0
		y0 = tempy1
		y1 = tempy0
	}
	dy = (float64(y1) - float64(y0))
	dx = (float64(x1) - float64(x0))

	offset := 0.0
	limit := 0.5

	
	m := (dy)/(dx)
	y := y0
	
	// fmt.Println("x0: ", x0, "x1: ", x1, "y0: ", y0, "y1: ", y1, "m", m)
	for x := x0; x <= x1; x++ {
		if steep {
			// Dibujar de manera vertical
			r.GlPoint(Point{y,x}, colors...)
		} else {
			// Dibujar de manera horizontal	
			r.GlPoint(Point{x,y}, colors...)
		}
		
		offset += m
		// Dibujar dependiendo si la linea va para abajo o para arriba.
		if math.Abs(offset) >= limit {
			if y0 < y1 {
				y += 1
			} else {
				y -= 1
			}
			limit += 1
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

// Draws a polygon with a set of points and a given color.
func (r * Renderer) GlPolygon(color Color,points ...Point) []Point {
	for i := 0; i < len(points); i++ {
		r.GLLine(points[i], points[(i+1) % len(points)], color)		
	}
	return points
}

// Fills a polygon with a set of points and a given color
func (r * Renderer) GlFillPolygon(color Color,points ...Point) []Point {
	filledPoints := []Point{}	// Points that were filled
	filledPoints = append(filledPoints, r.GlPolygon(color, points...)...) // Draw the polygon first
	xmin,xmax,ymin,ymax := glGetMaxMinDimensions(points)
	// Do a test for each point
	for y := int(ymin); y <= int(ymax); y++ {
		for x := int(xmin); x <= int(xmax); x++ {
			edgesRight := 0
			// edgesLeft := 0
			if (y == int(ymin) || x == int(xmin) )|| (y == int(ymax) || x == int(xmax)){
				continue
			}

			if glGetEdges(r.pixels[y], r.clearColor) <= 1 {
				continue
			}

			if (r.pixels[y][x] == color) {
				continue
			}

			if r.pixels[y][x] == color {
				continue
			}

			for j := x; j <= int(xmax); j++ {
				if r.pixels[y][j] == color && r.pixels[y][j + 1] != color {
					edgesRight += 1
				}
			}
			
			// for j := int(xmin); j < x; j++ {
			// 	if r.pixels[y][j] == color && r.pixels[y][j+1] != color {
			// 		edgesLeft += 1
			// 	}
			// }
			
			if  edgesRight % 2 != 0 {
				pointToFill := Point{float32(x),float32(y)}
				filledPoints = append(filledPoints, pointToFill)
				r.GlPoint(pointToFill, color)
			}
		}
		
	}
	return filledPoints
}

func glGetEdges(colors []Color, clearColor Color) int {
	edges := 0
	// If the array is empty, return 0
	if (len(colors) <= 0) {
		return edges
	}
	previousColor := colors[0]
	for i := 0; i < len(colors); i++ {
		// If the first pixel is already different, add it
		if i == 0 && colors[i] != clearColor {
			edges += 1
			continue
		}
		// Count the edges
		if colors[i] != previousColor && colors[i] != clearColor {
			edges += 1
		}
		previousColor = colors[i]	// Set the previous pixel
	}
	return edges
}

// Gets the minimum and maximum values for each axis for a set of points
func glGetMaxMinDimensions(points []Point) (float32, float32, float32, float32) {
	// Variables for dimensions
	ymin,ymax,xmin,xmax := points[0].Y, points[0].Y, points[0].X, points[0].X
	// Calculate the maximum and minimum values
	for i := 0; i < len(points); i++ {
		point := points[i]
		ymin = float32(math.Min(float64(ymin), float64(point.Y)))
		ymax = float32(math.Max(float64(ymax), float64(point.Y)))
		xmin = float32(math.Min(float64(xmin), float64(point.X)))
		xmax = float32(math.Max(float64(xmax), float64(point.X)))
	}
	return xmin,xmax,ymin,ymax
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