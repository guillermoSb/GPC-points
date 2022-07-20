package main

import (
	"fmt"
	"guillermoSb/glDots/gl"
)

func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	width := 960
	height := 540
	renderer := gl.Renderer{}
	renderer.GlInit(uint32(width), uint32(height))
	renderer.GlClearColor(0.1,0.1,0.1)
	renderer.GlColor(1,0,0)
	renderer.GlClear()
	
	poly1 := []gl.Point{{165, 380}, {185, 360}, 
						   {180, 330}, {207, 345}, 
						   {233, 330}, {230, 360}, 
						   {250, 380}, {220, 385}, 
						   {205, 410}, {193, 383}}

	
	renderer.GlPolygon(gl.Color{1,0,0}, poly1...)
	

	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
