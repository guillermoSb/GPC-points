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

func Test_pointIsInSlice_point_in_slice_returns_true(t *testing.T) {
	// Arrange
	points := []Point{{1,1}, {1,0}, {1,0}}
	// Act
	result := pointIsInSlice(points, Point{float32(1),float32(1)})
	assert.True(t, result)
}
func Test_pointIsInSlice_point_in_slice_returns_false(t *testing.T) {
	// Arrange
	points := []Point{{1,1}, {1,0}, {1,0}}
	// Act
	result := pointIsInSlice(points, Point{float32(2),float32(2)})
	assert.False(t, result)
}

func Test_glGetEdges_with_one_edge_returns_one(t *testing.T) {
	// Arrange
	pixels1 := []Color{{0,0,0},{1,0,0},{1,0,0},{0,0,0}}
	pixels2 := []Color{{0,0,0},{1,0,0},{0,0,0},{0,0,0}}
	pixels3 := []Color{{0,0,0},{0,0,0},{0,0,1}}
	pixels4 := []Color{{1,0,0},{0,0,0},{0,0,0}}
	pixels5 := []Color{{1,0,0},{0,0,0},{1,0,0}}
	pixels6 := []Color{{1,0,0},{1,0,0},{1,0,0}}
	// Act
	edges1 := glGetEdges(pixels1, Color{1,0,0})
	edges2 := glGetEdges(pixels2, Color{1,0,0})
	edges3 := glGetEdges(pixels3, Color{0,0,1})
	edges4 := glGetEdges(pixels4, Color{1,0,0})
	edges5 := glGetEdges(pixels5, Color{1,0,0})
	edges6 := glGetEdges(pixels6, Color{1,0,0})
	// Assert
	assert.Equal(t, 1, edges1)
	assert.Equal(t, 1, edges2)
	assert.Equal(t, 1, edges3)
	assert.Equal(t, 1, edges4)
	assert.Equal(t, 2, edges5)
	assert.Equal(t, 1, edges6)
}

func Test_glGetMaxMinDimensions_returns_accurate_information(t *testing.T) {
	// Arrange
	points := []Point{{5,4}, {10, 4}, {5,2}}
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