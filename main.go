package main

import (
	"fmt"
	"guillermoSb/glDots/gl"
	"math/rand"
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

	lines := 200
	offset := 0	
	for i := 0; i < lines; i++ {
		offset += width/lines
		r := rand.Float32()
		g := rand.Float32()
		b := rand.Float32()
		renderer.GLLine(gl.Point{float32(0), float32(height - offset)}, gl.Point{float32(offset), float32(0)}, gl.Color{r,g,b})
		renderer.GLLine(gl.Point{float32(width), float32(offset)}, gl.Point{float32(width - offset), float32(height)}, gl.Color{r,g,b})
	}

	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
