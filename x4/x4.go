package x4

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Header struct {
	mark string
	version uint16
	pageCount uint16
	readDirection uint8
	hasMetaData uint8
	hasThumbnails uint8
	hasChapters uint8
	currentPage uint32
	metadataOffset uint64
	indexOffset uint64
	dataOffset uint64
	thumbnailOffset uint64
	chapterOffset uint64
}

func getHeader(filePT *os.File) (Header, error){
	var header Header
	headerBuffer := make([]byte, 56)

	bufferReadLen, err := filePT.ReadAt(headerBuffer, 0)
	if err != nil && bufferReadLen != 56 {
		return fmt.Errorf("%v", err)
	}

	header.mark = string(headerBuffer[0:4])
	header.version = binary.LittleEndian.Uint16(headerBuffer[4:6])
	header.pageCount = binary.LittleEndian.Uint16(headerBuffer[6:8])
	header.readDirection = binary.LittleEndian.Uint8(headerBuffer[8:9])
	header.hasMetaData = binary.LittleEndian.Uint8(headerBuffer[9:10])
	header.hasThumbnails = binary.LittleEndian.Uint8(headerBuffer[10:11])
	header.hasChapters = binary.LittleEndian.Uint8(headerBuffer[11:12])
	header.currentPage = binary.LittleEndian.Uint32(headerBuffer[12:16])
	header.metadataOffset = binary.LittleEndian.Uint64(headerBuffer[16:24])
	header.indexOffset = binary.LittleEndian.Uint64(headerBuffer[24:32])
	header.dataOffset = binary.LittleEndian.Uint64(headerBuffer[32:40])
	header.thumbnailOffset = binary.LittleEndian.Uint64(headerBuffer[40:48])
	header.chapterOffset = binary.LittleEndian.Uint64(headerBuffer[48:56])

	return header, nil
}

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

func 