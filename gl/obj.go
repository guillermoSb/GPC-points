package gl

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Obj struct {
	Vertices []float32
	Texrecords []float32
	Normals []float32
	Faces [][]int
}

func (o *Obj) InitObj(filename string) Obj{
	obj := Obj{}
	dat, err := os.ReadFile(filename)
    if err != nil {
		fmt.Println(err)
		log.Fatal("Could not read file: ", filename)
	}
    lines := strings.Split(string(dat), "\n")
	// vertices := []float32{}
	// texrecords := []float32{}
	// normals := []float32{}
	// faces := [][]int{}
	for _, line := range lines {
		// Continue if len line is 0
		if len(line) == 0 {
			continue
		}
		prefix := strings.Split(line," ")[0]
		value := strings.Split(line, prefix)[1]
		
		switch prefix {
		case "v":
			fmt.Println(value)
		}
	}
	
	return obj
}

func stringSliceToFloat32Slice(slice []string) []float32{
	newSlice := []float32{}
	for _, v := range slice {
		if s,err := strconv.ParseFloat(v, 32); err == nil {
			newSlice = append(newSlice, float32(s))
		}
	}
	return newSlice
}

func stringSliceToIntArray(slice []string) []int{
	newSlice := []int{}
	for _, v := range slice {
		if s,err := strconv.ParseInt(v, 10, 32); err == nil {
			newSlice = append(newSlice, int(s))
		}
	}
	return newSlice
}