package gl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_can_draw_polygon(t *testing.T) {
	// Arrange
	renderer := createTestingRenderer(100, 100)
	polygon := []Point{{10,10}, {10, 20}, {20, 20}, {20,10}}
	// Act
	polyPoints := renderer.GlPolygon(Color{1,0,0}, polygon...)
	// Assert
	assert.NotEmpty(t,polyPoints)
}


func Test_can_fill_simple_polygon(t *testing.T) {
	// Arrange
	renderer := createTestingRenderer(100, 100)
	polygon := []Point{{2,2}, {2, 4}, {4, 4}, {4,2}}
	// Act
	filledPoints := renderer.GlFillPolygon(Color{1,0,0}, polygon...)
	// Assert
	assert.Len(t, filledPoints, 5)	// Expect 5 filled points returned.
}


func Test_can_fill_complex_polygon(t *testing.T) {
	// Arrange
	// renderer := createTestingRenderer(100, 100)
	// Act
	// Assert
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