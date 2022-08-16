package main

// turbosquid
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
	
	// PHOTO SHOOT

	// MEDIUM SHOT - TWO MARIOS
	// renderer.GlLoadModel("mario.obj", gl.V3{-3,0,-10}, gl.V3{0,0,0}, gl.V3{0.04,0.04,0.04})
	// renderer.GlLoadModel("mario.obj", gl.V3{1,-1,-3}, gl.V3{0,0,0}, gl.V3{0.04,0.04,0.04})
	// LOW ANGLE
	// renderer.GlLoadModel("mario.obj", gl.V3{0,0,-5}, gl.V3{-24,0,0}, gl.V3{0.04,0.04,0.04})
	// HIGH ANGLE
	// renderer.GlLoadModel("mario.obj", gl.V3{0,-1,-10}, gl.V3{25,0,0}, gl.V3{0.04,0.04,0.04})
	// DUTCH ANGLE
	renderer.GlLoadModel("mario.obj", gl.V3{0,0,-5}, gl.V3{0,0,45}, gl.V3{0.04,0.04,0.04})
	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
