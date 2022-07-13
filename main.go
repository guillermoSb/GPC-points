package main

import (
	"fmt"
	"guillermoSb/glDots/gl"
)

// ****************************************************************
// Main Program
// ****************************************************************


func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	width := 16
	height := 16
	renderer := gl.Renderer{}
	renderer.GlInit(uint32(width), uint32(height))
	renderer.GlClearColor(0.1,0.1,0.1)
	renderer.GlColor(1,0,0)
	renderer.GlClear()
	renderer.GlViewPort(0,0,1,1)
	renderer.GlClearViewport(gl.Color{1,1,1})
	renderer.GlColor(0.5,0.5,0.5)
	renderer.GlViewPortPoint(gl.Point{1,1})
	renderer.GlViewPortPoint(gl.Point{0,0})
	renderer.GlViewPortPoint(gl.Point{-1,-1})
	// renderer.GlViewPortPoint(gl.Point{0,0})


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
