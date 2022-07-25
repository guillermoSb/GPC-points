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
	renderer.GlColor(1,1,1)
	renderer.GlClear()

	// Object
	
	renderer.GlLoadModel("model.obj", gl.V3{float32(width)/2,float32(height)/2,50}, gl.V3{50,50,50}, gl.V3{250,250,200})
	

	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
