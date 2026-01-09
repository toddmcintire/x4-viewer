package x4

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

//56 bytes
type Header struct {
	mark string
	version uint16
	pageCount uint16
	readDirection ReadDirection
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

//256 bytes (optional, at metadataOffset)
type Metadata struct {
	title [128]byte
	author [64]byte
	publisher [32]byte
	language [16]byte
	createTime uint32
	coverPage uint16
	chapterCount uint16
	reserved uint64
}

// n * 96 bytes (optional, at chapterOffset)
type Chapter struct {
	chapterName [80]byte
	startPage uint16
	endPage uint16
	reserved1 uint32
	reserved2 uint32
	reserved3 uint32
}

// pageCount * 16 bytes (at indexOffset)
//thumbnail area (optional at thumbOffset after pageData)
type Page struct {
	offset uint64
	size uint32
	width uint16
	height uint16
	dataOffset [48000]byte
}

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
	header.readDirection = ReadDirection(uint8(binary.LittleEndian.Uint16(headerBuffer[8:9])))
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