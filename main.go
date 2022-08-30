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
	renderer.GlInit(uint32(width), uint32(height), "sunshine-bg.bmp")
	renderer.GlClearColor(0,0,0)
	renderer.GlColor(1,1,1)
	renderer.GlClear()
	renderer.GLDrawBackground()
	
	// PHOTO SHOOT
	// Para utilizar el photo shoot descompentar las lineas de cada una
	// MEDIUM SHOT - TWO MARIOS
	// renderer.GlLoadModel("mario.obj", gl.V3{-3,0,-10}, gl.V3{0,0,0}, gl.V3{0.04,0.04,0.04})
	// renderer.GlLoadModel("mario.obj", gl.V3{1,-1,-3}, gl.V3{0,0,0}, gl.V3{0.04,0.04,0.04})
	// LOW ANGLE
	// renderer.GlLoadModel("mario.obj", gl.V3{0,0,-5}, gl.V3{-24,0,0}, gl.V3{0.04,0.04,0.04})
	// HIGH ANGLE
	// renderer.GlLoadModel("mario.obj", gl.V3{0,-1,-10}, gl.V3{25,0,0}, gl.V3{0.04,0.04,0.04})
	// DUTCH ANGLE

	// renderer.ActiveShader = gl.UnlitShader
	// renderer.GlLoadModel("lucky.obj", gl.V3{0,0,-4}, gl.V3{0,0,0}, gl.V3{0.00005,0.00005,0.00005})
	// renderer.ActiveTexture = gl.InitTexture("marioD.bmp")	// Mario Texture
	// renderer.ActiveShader = gl.FlatShader
	// renderer.GlLoadModel("mario.obj", gl.V3{1,-1,-12}, gl.V3{0,0,0}, gl.V3{0.02,0.02,0.02})
	// renderer.ActiveTexture = gl.InitTexture("plane.bmp")	// Mario Texture
	renderer.GlLoadModel("plane.obj", gl.V3{0,0,-100}, gl.V3{0,0,0}, gl.V3{0.02,0.02,0.02})
	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
