package gl

import (
	"encoding/binary"
	"guillermoSb/glLibrary/numg"
	"log"
	"math"
	"math/rand"
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
func (r * Renderer) GlPoint(point V2, colors ...Color) {
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
func (r * Renderer) GlViewPortPoint(point V2) {
	if point.X < -1 || point.X > 1 || point.Y < -1 || point.Y > 1 {
		return;
	}
	x := (point.X + 1) * ((float32(r.vpWidth - 1) / 2 )) + float32(r.vpX)
	y := (point.Y + 1) * ((float32(r.vpHeight - 1) / 2 )) + float32(r.vpY)
	r.GlPoint(V2{x,y})
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
func (r *Renderer) GLLine(v0, v1 V2, colors ...Color) {
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
			r.GlPoint(V2{y,x}, colors...)
		} else {
			// Dibujar de manera horizontal	
			r.GlPoint(V2{x,y}, colors...)
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

func (r * Renderer) GlFillPolygon(color Color,points ...V2) []V2  {
	r.GlPolygon(color, points...)	// Create the polygon first
	xmin,xmax,ymin,ymax := glGetMaxMinDimensions(points)	// Get the min and max dimensions
	// Get a middle point
	middlePoint := V2{float32(int(xmax - (xmax-xmin) / 2)), float32(int(ymax - (ymax-ymin) / 2)),}
	// Get the first color
	oldColor := r.pixels[int(middlePoint.Y)][int(middlePoint.X)]
	// Do not draw if the color is already there.
	if oldColor == color {
		return []V2{}
	}
	filledPoints := r.fillPoint(middlePoint, color, oldColor)
	return filledPoints
}

func (r * Renderer) GlFillPolygonCrazy(color Color,points ...V2) []V2  {
	filledPoints := r.GlFillPolygon(color,points...)
	trials := 100
	for i := 0; i < trials; i++ {
		idx := rand.Intn(len(filledPoints))
		top := V2{filledPoints[idx].X, filledPoints[idx].Y + 1}
		bottom := V2{filledPoints[idx].X, filledPoints[idx].Y - 1}
		left := V2{filledPoints[idx].X + 1, filledPoints[idx].Y}
		right := V2{filledPoints[idx].X - 1, filledPoints[idx].Y}

		red := rand.Float32()
		green := rand.Float32()
		blue := rand.Float32()
	
		r.GlPoint(top, Color{red,green,blue})
		r.GlPoint(right, Color{red,green,blue})
		r.GlPoint(bottom, Color{red,green,blue})
		r.GlPoint(left, Color{red,green,blue})
		r.GlPoint(filledPoints[idx], Color{red,green,blue})
		
		
	}
	return filledPoints
}



func (r * Renderer) fillPoint(point V2, color Color, oldColor Color) []V2 {
	// Return if the point is already filled
	m := r.width
	n := r.height
	points := []V2{}
	if !(point.X < 0 || point.X > float32(m) || point.Y < 0 || point.Y > float32(n) || r.pixels[int(point.Y)][int(point.X)] != oldColor) {
		r.GlPoint(point, color)
		points = append(points, point)
		top := V2{point.X, point.Y + 1}
		bottom := V2{point.X, point.Y - 1}
		left := V2{point.X + 1, point.Y}
		right := V2{point.X - 1, point.Y}
		points = append(points, r.fillPoint(bottom,color, oldColor)...)
		points = append(points, r.fillPoint(top,color, oldColor)...)
		points = append(points, r.fillPoint(right,color, oldColor)...)
		points = append(points, r.fillPoint(left,color, oldColor)...)
	}
	return points;
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
func (r * Renderer) GlPolygon(color Color,points ...V2) []V2 {
	for i := 0; i < len(points); i++ {
		r.GLLine(points[i], points[(i+1) % len(points)], color)		
	}
	return points
}

func (r * Renderer) GlLoadModel(filename string,translate, rotate, scale V3) {
	model := Obj{}
	model = model.InitObj(filename)
	modelMatrix := glCreateObjectMatrix(translate, rotate, scale)
	for _, face := range model.Faces {
		vertCount := len(face)
		for i := 0; i < vertCount; i++ {
			v0 := model.Vertices[face[i][0] - 1]
			v1 := model.Vertices[face[(i + 1) % vertCount][0] - 1]
			p1 := glTransform(V3{v0[0], v0[1], v0[2]}, modelMatrix)
			p2 := glTransform(V3{v1[0], v1[1], v1[2]}, modelMatrix)
			
			r.GLLine(V2{p1.X, p1.Y}, V2{p2.X, p2.Y}, Color{rand.Float32(), rand.Float32(), rand.Float32()})	
		}
	}

}

func glTransform(vertex V3, matrix numg.M) V3 {
	v := V4{vertex.X, vertex.Y, vertex.Z,1}
	vt, _ := numg.MultiplyMatrices(matrix, numg.M{{v.X}, {v.Y}, {v.Z},{v.W}})
	vf := V3{vt[0][0]/vt[3][0], vt[1][0]/vt[3][0], vt[2][0]/vt[3][0]}
	return vf
}

func glCreateObjectMatrix(translate, rotate, scale V3) numg.M {
	translateMatrix, _ := numg.Identity(4)
	translateMatrix[0][3] = translate.X
	translateMatrix[1][3] = translate.Y
	translateMatrix[2][3] = translate.Z

	scaleMatrix, _ := numg.Identity(4)
	scaleMatrix[0][0] = scale.X
	scaleMatrix[1][1] = scale.Y
	scaleMatrix[2][2] = scale.Z

	result, _ := numg.MultiplyMatrices(translateMatrix, scaleMatrix)
	return result

}


func glGetEdges(colors []Color, color Color) int {
	edges := 0
	// If the array is empty, return 0
	if (len(colors) <= 0) {
		return edges
	}
	previousColor := colors[0]
	for i := 0; i < len(colors); i++ {
		// If the first pixel is already different, add it
		if i == 0 && colors[i] == color {
			edges += 1
			continue
		}
		// Count the edges
		if colors[i] != previousColor && colors[i] == color {
			
			edges += 1
		}
		previousColor = colors[i]	// Set the previous pixel
	}
	return edges
}

// Gets the minimum and maximum values for each axis for a set of points
func glGetMaxMinDimensions(points []V2) (float32, float32, float32, float32) {
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

// Finds if a point is inside a slice of points
func pointIsInSlice(points []V2, point V2) bool {
	for _, v := range points {
		if v == point {
			return true
		}
	}
	return false
}



// ****************************************************************
// Utils and Structures
// ****************************************************************

type V2 struct {
	X float32
	Y float32
}

type V3 struct {
	X float32
	Y float32
	Z float32
}

type V4 struct {
	X float32
	Y float32
	Z float32
	W float32
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