package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	bs := make([]byte,4)
	binary.LittleEndian.PutUint32(bs, 32)
	fmt.Println(bs)
}