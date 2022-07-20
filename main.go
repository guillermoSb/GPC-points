package main

import (
	"fmt"
	"guillermoSb/glDots/gl"
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

	renderer.GlPolygon(gl.Color{1,0,0}, gl.Point{50, 50}, gl.Point{50, 200},gl.Point{50, 200}, gl.Point{200, 200},gl.Point{200, 200}, gl.Point{200, 50},gl.Point{200, 50}, gl.Point{50, 50})

	// Triangle
	renderer.GlPolygon(gl.Color{1,0,0}, gl.Point{300, 50}, gl.Point{500, 50}, gl.Point{300, 50}, gl.Point{400, 200}, gl.Point{500, 50}, gl.Point{400, 200})

	// Triangle 90 degrees
	renderer.GlPolygon(gl.Color{1,0,0},gl.Point{600,50}, gl.Point{750,50}, gl.Point{600,50},gl.Point{600,200}, gl.Point{600,200}, gl.Point{750,50})

	// Arrow
	renderer.GlPolygon(gl.Color{0,1,0}, gl.Point{50,400}, gl.Point{50,450}, gl.Point{150,450}, gl.Point{150, 475}, gl.Point{200,425},gl.Point{150, 375}, gl.Point{150, 400})

	// Hexagon
	renderer.GlPolygon(gl.Color{0,0,1}, gl.Point{300, 425}, gl.Point{375, 375}, gl.Point{450, 425}, gl.Point{450, 500},
					   gl.Point{375,550}, gl.Point{300, 500})
	// Pentagon
	renderer.GlPolygon(gl.Color{0,1,1}, gl.Point{600, 375}, gl.Point{750, 375}, gl.Point{780, 500}, gl.Point{675, 550}, gl.Point{575, 500})
	// Hexagon
	// angle := math.Pi - (2/3) * math.Pi
	// yprev := 50.0
	// xprev := 600.0
	// for i := 0; i < 6; i++ {
	// 	renderer.GLLine((gl.Point{float32(xprev), float32(yprev)}), gl.Point{float32(math.Cos(200 * angle)), float32(math.Sin(200 * angle))})
	// }



	
	

	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
