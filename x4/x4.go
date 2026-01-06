package x4

import (
	"fmt"
	"os"
)

//given path will return a slice of bytes
func GetXTGData(path string, buf []byte) int {
	fmt.Println(path)
	filePtr, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(filePtr)

	i, err :=filePtr.ReadAt(buf, 22)	
	if err != nil {
		panic(err)
	}
	return i
}