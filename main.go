package main

import (
	"fmt"
	"guillermoSb/glDots/gl"
	"math"
)

func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	width := 1024
	height := 1024
	renderer := gl.Renderer{}
	renderer.GlInit(uint32(width), uint32(height))
	renderer.GlClearColor(0.1,0.1,0.1)
	renderer.GlColor(1,0,0)
	renderer.GlClear()

	// Square:
	renderer.GLLine(gl.Point{50, 50}, gl.Point{50, 200}) // |
	renderer.GLLine(gl.Point{50, 200}, gl.Point{200, 200}) // -
	renderer.GLLine(gl.Point{200, 200}, gl.Point{200, 50}) // |
	renderer.GLLine(gl.Point{200, 50}, gl.Point{50, 50}) // -

	// Triangle

	renderer.GLLine(gl.Point{300, 50}, gl.Point{500, 50}) // -
	renderer.GLLine(gl.Point{300, 50}, gl.Point{400, 200}) // -
	renderer.GLLine(gl.Point{500, 50}, gl.Point{400, 200}) // -

	// Hexagon
	angle := math.Pi - (2/3) * math.Pi
	yprev := 50.0
	xprev := 600.0
	for i := 0; i < 6; i++ {
		renderer.GLLine((gl.Point{float32(xprev), float32(yprev)}), gl.Point{float32(math.Cos(200 * angle)), float32(math.Sin(200 * angle))})
	}



	
	

	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
