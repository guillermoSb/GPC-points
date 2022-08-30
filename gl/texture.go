package gl

import (
	"encoding/binary"
	"errors"
	"log"
	"os"
)

type Texture struct {
	Width, Height int
	Pixels [][]Color
	Name string
}

func InitTexture(fileName string) Texture{
	texture := Texture{}
	f, err := os.Open(fileName)

	// Check if the file was successfully created
	if err != nil {
		log.Fatal(err)
	}
	// Example, writing 5 x 5 image
	defer f.Close()	// Close the file when the process is done
	headerSizeBits := make([]byte, 4)
  f.ReadAt(headerSizeBits, 10)
	headerSize := fromByteToInt(headerSizeBits)
	
	widthBits := make([]byte, 4)
	f.ReadAt(widthBits, 18)
	width := fromByteToInt(widthBits)

	heightBits := make([]byte, 4)
	f.ReadAt(heightBits, 22)
	height := fromByteToInt(heightBits)

	f.Seek(int64(headerSize), 0)
	pixels := [][]Color{}

	for y := 0; y < int(height); y++ {
		row := []Color{}
		for x := 0; x < int(width); x++ {
			bByte := make([]byte,1)
			gByte := make([]byte,1)
			rByte := make([]byte,1)
			f.Read(bByte)
			f.Read(gByte)
			f.Read(rByte)
			// r := fromByteToInt(bByte)
			r := float32(float32(rByte[0]) / 255.0)
			g := float32(float32(gByte[0]) / 255.0)
			b := float32(float32(bByte[0]) / 255.0)
			row = append(row, Color{r,g,b})
		}
		pixels = append(pixels, row)
	}
  texture.Width = int(width)
	texture.Height = int(height)
	// 🔥 quemando la textura
	texture.Pixels = pixels
	texture.Name = fileName
	return texture
}

func (t *Texture) GetColor(u,v float32) (Color, error) {

	if 0 <= u && u <= 1 && 0 <= v && v <= 1{
		return t.Pixels[int((v) * float32(t.Height))][int((u) * float32(t.Width))], nil
	} 
	return Color{}, errors.New("Cannot get the color")
}

func fromByteToInt(bytes []byte) uint32 {
	return binary.LittleEndian.Uint32(bytes)
}