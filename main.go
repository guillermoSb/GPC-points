package main

import (
	"fmt"
	"guillermoSb/glLibrary/gl"
)

func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	width := 1024
	height := 1024
	renderer := gl.Renderer{}
	renderer.GlInit(uint32(width), uint32(height))
	renderer.GlClearColor(0.1,0.1,0.1)
	renderer.GlColor(1,1,1)
	renderer.GlClear()
	renderer.UseShader = true
	// Object
	renderer.GlLoadModel("mario.obj", gl.V3{float32(width)/2,float32(height)/2,0}, gl.V3{-30,180,80}, gl.V3{5,5,5})

	// renderer.GlLoadModel("sphere.obj", gl.V3{float32(width)/2,0,1}, gl.V3{50,50,50}, gl.V3{500,500,0})
	// renderer.GlLoadModel("sphere.obj", gl.V3{float32(width),float32(height),1}, gl.V3{50,50,50}, gl.V3{250,250,200})
	// renderer.GlLoadModel("sphere.obj", gl.V3{float32(0),float32(0),1}, gl.V3{50,50,50}, gl.V3{250,250,200})
	// Triangle
	
	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
