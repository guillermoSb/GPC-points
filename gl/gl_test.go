package gl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_can_draw_polygon(t *testing.T) {
	// Arrange
	renderer := createTestingRenderer(100, 100)
	polygon := []V2{{10,10}, {10, 20}, {20, 20}, {20,10}}
	// Act
	polyPoints := renderer.GlPolygon(Color{1,0,0}, polygon...)
	// Assert
	assert.NotEmpty(t,polyPoints)
}


func Test_can_fill_simple_polygon(t *testing.T) {
	// Arrange
	renderer := createTestingRenderer(100, 100)
	polygon := []V2{{2,2}, {2, 4}, {5, 4}, {5,2}}
	// Act
	fmt.Println(len(renderer.GlFillPolygon(Color{1,0,0}, polygon...)))
	// Assert
	assert.Equal(t,Color{1,0,0}, renderer.pixels[3][3])
	assert.Equal(t,Color{1,0,0}, renderer.pixels[3][4])
	
}


func Test_glGetMaxMinDimensions_returns_accurate_information(t *testing.T) {
	// Arrange
	points := []V2{{5,4}, {10, 4}, {5,2}}
	// Act
	minX,maxX,minY,maxY := glGetMaxMinDimensions(points)
	// Assert
	assert.Equal(t, float32(2.0), minY)
	assert.Equal(t, float32(4.0), maxY)
	assert.Equal(t, float32(5.0), minX)
	assert.Equal(t, float32(10.0), maxX)

}



func createTestingRenderer(vWidth, vHeight int) Renderer {
	width := vWidth
	height := vHeight
	renderer := Renderer{}
	renderer.GlInit(uint32(width), uint32(height))
	renderer.GlClearColor(0.1,0.1,0.1)
	renderer.GlColor(1,0,0)
	renderer.GlClear()
	return renderer
}