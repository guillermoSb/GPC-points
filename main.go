package main

import (
	"fmt"
	"guillermoSb/glLibrary/gl"
)

func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	width := 800
	height := 800
	renderer := gl.Renderer{}
	renderer.GlInit(uint32(width), uint32(height))
	renderer.GlClearColor(0.1,0.1,0.1)
	renderer.GlColor(1,0,0)
	renderer.GlClear()
	
	// Define the polygon dots
	poly1 := []gl.V2{{165, 380}, {185, 360}, 
						   {180, 330}, {207, 345}, 
						   {233, 330}, {230, 360}, 
						   {250, 380}, {220, 385}, 
						   {205, 410}, {193, 383}}

	poly2 := []gl.V2{{321, 335}, {288, 286}, {339, 251}, {374, 302}}
	poly3 := []gl.V2{{377, 249}, {411, 197}, {436, 249}}
	poly4 := []gl.V2{{413, 177}, {448, 159}, {502, 88}, {553, 53}, {535, 36}, {676, 37}, {660, 52},
						{750, 145}, {761, 179}, {672, 192}, {659, 214}, {615, 214}, {632, 230}, {580, 230},
						{597, 215}, {552, 214}, {517, 144}, {466, 180},}
	poly5 := []gl.V2{{682, 175}, {708, 120}, {735, 148}, {739, 170}}
	
	// Draw polygons

	renderer.GlFillPolygon(gl.Color{0,1,0}, poly1...)
	renderer.GlFillPolygon(gl.Color{0,0,1}, poly2...)
	renderer.GlFillPolygon(gl.Color{1,0,0}, poly3...)
	renderer.GlFillPolygon(gl.Color{1,1,1}, poly4...)
	renderer.GlFillPolygon(gl.Color{1,0,1}, poly5...)
	// renderer.GlFillPolygon(gl.Color{1,1,0}, poly5...)


	

	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
