package main

import (
	"fmt"
	"guillermoSb/glDots/gl"
)

// ****************************************************************
// Main Program
// ****************************************************************


func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	renderer := gl.Renderer{}
	renderer.GlCreateWindow(40,40);
	renderer.GlClearColor(0.5,0,0)
	renderer.GlClear()
	renderer.GlColor(0,1,0)
	renderer.GlPoint(gl.Point{X: 0.75,Y: 0.75})
	renderer.GlColor(0,0,1)
	renderer.GlPoint(gl.Point{X: 1,Y: 0.99})
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
