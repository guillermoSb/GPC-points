package main // Main package
import (
	"fmt"
	"log"
	"os"
)

// Encoding binary
//

func main() {
	fmt.Println("---------- BMP IMAGE CREATION ----------")	
	renderer := Renderer{width: 10, height: 10}
	renderer.glFinish("out.bmp")
}



type Renderer struct {
	width, height uint32
}


func (r * Renderer) glFinish(fileName string) {
	// Attempt to open the file
	f, err := os.Create(fileName)

	// Check if the file was successfully created
	if err != nil {
		log.Fatal(err)
	}
	// Example, writing 5 x 5 image
	defer f.Close()	// Close the file when the process is done
	f.Write([]byte("B"))
	f.Write([]byte("M"))
	f.Write([]byte{0, 0, 0, 0})	// File Size
	f.Write([]byte{0, 0})	// Reserved
	f.Write([]byte{0, 0})	// Reserved
	f.Write([]byte{54, 0, 0, 0 })	// ?
	f.Write([]byte{40, 0, 0, 0})	// Header Size
	f.Write([]byte{4, 0, 0, 0})		// Width
	f.Write([]byte{4, 0, 0, 0})		// Height
	f.Write([]byte{1, 0})	// Plane
	f.Write([]byte{24, 0})	// BPP
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	// Pixel Data

	f.Write([]byte{107, 231, 59})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})

	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})

	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})

	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})

	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})
	f.Write([]byte{0,0,0})

	f.Write([]byte{0,0,0,0,0})






	// fmt.Println(binary.BigEndian.Uint16([]byte("BM")))
}

type Color struct {
	r, g, b uint8
}





