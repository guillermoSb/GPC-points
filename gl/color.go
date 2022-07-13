package gl

import (
	"fmt"
)

type Color struct {
	R, G, B float32
}

func (c * Color ) bytes()[]byte {
	red := uint8(c.R * 255)
	green := uint8(c.G * 255)
	blue := uint8(c.B * 255)
	return []byte{blue, green, red}
}

func ColorFromRGB(r,g,b float32) (Color, error) {
	// Check if the colors are valid
	if r > 1 || g > 1 || b > 1 || r < 0 || b < 0 || g < 0{
		return Color{}, fmt.Errorf("Can't create color from the received values. Expected r=%.1f, g=%.1f, b=%.1f to be within 0 and 1.", r,g,b)
	}
	return Color{r,g,b}, nil
}