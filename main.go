package main

import (
	"fmt"
	"guillermoSb/glLibrary/gl"
	"math/rand"
)

func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	width := 400
	height := 400
	renderer := gl.Renderer{}
	renderer.GlInit(uint32(width), uint32(height))
	renderer.GlClearColor(0.1,0.1,0.1)
	renderer.GlColor(1,1,1)
	renderer.GlClear()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if rand.Float32() > 0.98 {
				brightness := rand.Float32() / 2 + 0.1
				size := rand.Intn(3)
				if size == 0 {
					renderer.GlColor(0.3,0.3, 0.3)
					renderer.GlPoint(gl.V2{float32(x),float32(y)})
				} else if size == 1 {
					renderer.GlColor(0.3, 0.3, 0.3)
					renderer.GlPoint(gl.V2{float32(x + 1),float32(y)})
					renderer.GlPoint(gl.V2{float32(x + 1),float32(y + 1)})
					renderer.GlPoint(gl.V2{float32(x),float32(y + 1)})
				} else if size == 2 {
					renderer.GlColor(brightness, brightness, brightness)
					renderer.GlPoint(gl.V2{float32(x),float32(y + 1)})
					renderer.GlPoint(gl.V2{float32(x),float32(y)})
					renderer.GlPoint(gl.V2{float32(x + 1),float32(y)})
					renderer.GlPoint(gl.V2{float32(x - 1),float32(y)})
					renderer.GlPoint(gl.V2{float32(x),float32(y - 1)})
					
				}
			}
		}
	}
	// Object
	renderer.GlLoadModel("model.obj", gl.V3{float32(width)/2,float32(height)/2,1}, gl.V3{50,50,50}, gl.V3{100,100,0})
	// renderer.GlLoadModel("sphere.obj", gl.V3{float32(width)/2,0,1}, gl.V3{50,50,50}, gl.V3{500,500,0})
	// renderer.GlLoadModel("sphere.obj", gl.V3{float32(width),float32(height),1}, gl.V3{50,50,50}, gl.V3{250,250,200})
	// renderer.GlLoadModel("sphere.obj", gl.V3{float32(0),float32(0),1}, gl.V3{50,50,50}, gl.V3{250,250,200})
	// Triangle
	
	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
