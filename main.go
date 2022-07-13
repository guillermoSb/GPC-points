package main

import (
	"fmt"
	"guillermoSb/glDots/gl"
	"math/rand"
)

func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	width := 100
	height := 100
	renderer := gl.Renderer{}
	renderer.GlInit(uint32(width), uint32(height))
	renderer.GlClearColor(0.1,0.1,0.1)
	renderer.GlColor(1,0,0)
	renderer.GlClear()
	renderer.GlViewPort(100/2 - 25,100/2 - 25,50,50)
	
	// Creating a galaxy on the entire screen
	for x := 0; x < 960; x++ {
		for y := 0; y < 540; y++ {
			if rand.Float32() > 0.98 {
				brightness := rand.Float32() / 2 + 0.5
				size := rand.Intn(3)
				if size == 0 {
					renderer.GlColor(1, 0, 0)
					renderer.GlPoint(gl.Point{float32(x),float32(y)})
				} else if size == 1 {
					renderer.GlColor(1, 1, 0)
					renderer.GlPoint(gl.Point{float32(x + 1),float32(y)})
					renderer.GlPoint(gl.Point{float32(x + 1),float32(y + 1)})
					renderer.GlPoint(gl.Point{float32(x),float32(y + 1)})
				} else if size == 2 {
					renderer.GlColor(0, 0, brightness)
					renderer.GlPoint(gl.Point{float32(x),float32(y + 1)})
					renderer.GlPoint(gl.Point{float32(x),float32(y)})
					renderer.GlPoint(gl.Point{float32(x + 1),float32(y)})
					renderer.GlPoint(gl.Point{float32(x - 1),float32(y)})
					renderer.GlPoint(gl.Point{float32(x),float32(y - 1)})
					
				}
			}
		}
	}

	renderer.GlClearViewport(gl.Color{1,1,1})
	
	
	// Create Random Pattern on the viewport
	renderer.GlColor(0.5,0.5,0.5)
	for x := -1.0; x < 1; x += 0.1 {
		for y := -1.0; y < 1; y += 0.1 {
			thr := rand.Float32() 
			if thr > 0.2 {
				renderer.GlViewPortPoint(gl.Point{float32(x), float32(y)})
			}
		}
	}


	// EJEMPLOS EN CLASE

	// Dibujar estatica
	// for x := 0; x < 960; x++ {
	// 	for y := 0; y < 540; y++ {
	// 		renderer.GlPoint(gl.Point{rand.Float32(),rand.Float32()})
	// 	}
	// }

	// Colores Random
	// for x := 0; x < 960; x++ {
	// 	for y := 0; y < 540; y++ {
	// 		renderer.GlColor(rand.Float32(),rand.Float32(),rand.Float32())
	// 		renderer.GlPoint(gl.Point{float32(x),float32(y)})
	// 	}
	// }

	// Dibujando linea
	// for x := 0; x < 960; x++ {
	// 	renderer.GlPoint(gl.Point{float32(x),0})		
	// }

	// Ejemplo Galaxias
	// for x := 0; x < 960; x++ {
	// 	for y := 0; y < 540; y++ {
	// 		if rand.Float32() > 0.98 {
	// 			brightness := rand.Float32() / 2 + 0.5
	// 			size := rand.Intn(3)
	// 			if size == 0 {
	// 				renderer.GlColor(1, 0, 0)
	// 				renderer.GlPoint(gl.Point{float32(x),float32(y)})
	// 			} else if size == 1 {
	// 				renderer.GlColor(1, 1, 0)
	// 				renderer.GlPoint(gl.Point{float32(x + 1),float32(y)})
	// 				renderer.GlPoint(gl.Point{float32(x + 1),float32(y + 1)})
	// 				renderer.GlPoint(gl.Point{float32(x),float32(y + 1)})
	// 			} else if size == 2 {
	// 				renderer.GlColor(brightness, brightness, brightness)
	// 				renderer.GlPoint(gl.Point{float32(x),float32(y + 1)})
	// 				renderer.GlPoint(gl.Point{float32(x),float32(y)})
	// 				renderer.GlPoint(gl.Point{float32(x + 1),float32(y)})
	// 				renderer.GlPoint(gl.Point{float32(x - 1),float32(y)})
	// 				renderer.GlPoint(gl.Point{float32(x),float32(y - 1)})
					
	// 			}
	// 		}
	// 	}
	// }
	
	renderer.GlFinish("out.bmp")
	fmt.Println("----------      FINISHED      ----------")	
}
