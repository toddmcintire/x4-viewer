package x4

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type Header struct {
	mark string
	version uint16
	pageCount uint16
	ReadDirection uint8
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

type ReadDirection uint8

const (
	LR ReadDirection = iota
	RL
	TB
)

func GetXTCHeader(path string) (Header, error){
	filePT, openErr := os.Open(path)
	if openErr != nil {
		panic("error opening file")
	}
	var header Header
	headerBuffer := make([]byte, 56)

	bufferReadLen, err := filePT.ReadAt(headerBuffer, 0)
	if err != nil && bufferReadLen != 56 {
		return Header{}, fmt.Errorf("%v", err)
	}

	header.mark = string(headerBuffer[0:4])
	header.version = binary.LittleEndian.Uint16(headerBuffer[4:6])
	header.pageCount = binary.LittleEndian.Uint16(headerBuffer[6:8])
	header.readDirection = uint8(binary.LittleEndian.Uint16(headerBuffer[8:9]))
	header.hasMetaData = uint8(binary.LittleEndian.Uint16(headerBuffer[9:10]))
	header.hasThumbnails = uint8(binary.LittleEndian.Uint16(headerBuffer[10:11]))
	header.hasChapters = uint8(binary.LittleEndian.Uint16(headerBuffer[11:12]))
	header.currentPage = binary.LittleEndian.Uint32(headerBuffer[12:16])
	header.metadataOffset = binary.LittleEndian.Uint64(headerBuffer[16:24])
	header.indexOffset = binary.LittleEndian.Uint64(headerBuffer[24:32])
	header.dataOffset = binary.LittleEndian.Uint64(headerBuffer[32:40])
	header.thumbnailOffset = binary.LittleEndian.Uint64(headerBuffer[40:48])
	header.chapterOffset = binary.LittleEndian.Uint64(headerBuffer[48:56])
	
	if header.hasMetaData == 0 {
		return Header{}, errors.New("no metadata")
	}

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