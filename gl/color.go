package gl

import (
	"fmt"
)

// Structure that defines the properties of a Color.
// R: Value of the red component. Must be between 0 and 1
// G: Value of the green component. Must be between 0 and 1
// B: Value of the blue component. Must be between 0 and 1
type Color struct {
	R, G, B float32
}

// bytes allows to get the value of a color represented as a byte array of three components.
// Parameters:
// - c: Pointer to the Color
func (c * Color ) bytes()[]byte {
	red := uint8(c.R * 255)
	green := uint8(c.G * 255)
	blue := uint8(c.B * 255)
	return []byte{blue, green, red}
}

// ColorFromRGB converts an r,g,b value and returns a Color struct.
// Parameters:
// - r: Red component between 0 and 255
// - g: Green component between 0 and 255
// - b: Blue component between 0 and 255
func ColorFromRGB(r,g,b float32) (Color, error) {
	// Check if the colors are valid
	if r > 1 || g > 1 || b > 1 || r < 0 || b < 0 || g < 0{
		return Color{}, fmt.Errorf("Can't create color from the received values. Expected r=%.1f, g=%.1f, b=%.1f to be within 0 and 1.", r,g,b)
	}
	return Color{r,g,b}, nil
}