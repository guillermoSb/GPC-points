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
	

	// Mario
	renderer.ActiveTexture = gl.InitTexture("marioD.bmp")	// Mario Texture
	renderer.ActiveShader = gl.GShader
	renderer.GlLoadModel("mario.obj", gl.V3{1,-1,-12}, gl.V3{0,-30,0}, gl.V3{0.02,0.02,0.02})

	// Duck
	renderer.ActiveTexture = gl.InitTexture("marioD.bmp")	// Mario Texture
	renderer.ActiveShader = gl.YellowShader
	renderer.GlLoadModel("duck.obj", gl.V3{-1.2, -8, -12}, gl.V3{0,30,0}, gl.V3{0.002,0.002,0.002})
	// Plane
	renderer.ActiveTexture = gl.InitTexture("marioD.bmp")
	renderer.ActiveShader = gl.CShader
	renderer.GlLoadModel("plane.obj", gl.V3{1,4,-12}, gl.V3{0,30,0}, gl.V3{0.008,0.008,0.008 })
	// Hydrant
	renderer.ActiveTexture = gl.InitTexture("hydrant.bmp")
	renderer.ActiveShader = gl.BWShader
	renderer.GlLoadModel("hydrant.obj", gl.V3{0.5,-0.5,-2}, gl.V3{0,30,0}, gl.V3{0.004,0.004,0.004 })

	// Rock
	renderer.ActiveTexture = gl.InitTexture("rock.bmp")
	renderer.ActiveShader = gl.FlatShader
	renderer.GlLoadModel("rock.obj", gl.V3{0,-2,-10}, gl.V3{0,30,0}, gl.V3{0.004,0.004,0.004 })

	
	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
