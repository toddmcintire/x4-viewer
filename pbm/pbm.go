package pbm

import (
	//"strings"
	"fmt"
	"os"
	"strconv"
)

func CreatePBM(width int, height int, data []byte) error {
	createdFile, err := os.Create("page1.pbm")
	if err != nil {
		return fmt.Errorf("failed to create file %v", err)
	}
	createdFile.Write([]byte{'4', '\n'})

	
	createdFile.Write([]byte(strconv.Itoa(width)))
	createdFile.Write([]byte(" "))
	createdFile.Write([]byte(strconv.Itoa(height)))	
	createdFile.Write([]byte("\n"))

	//TODO loop through data 60 hex at a time (480)

	return nil
}

func TestLoop(data []byte) {
    chunkSize := 60
    
    for i := 0; i < len(data); i += chunkSize {
        // Calculate end index (don't go past array bounds)
        end := i + chunkSize
        if end > len(data) {
            end = len(data)
        }
        
        // Get chunk and process it
        chunk := data[i:end]
        
        //fmt.Printf("Processing chunk %d-%d: %v\n", i, end-1, chunk)
        // Your processing logic here
		var str string
		for _, v := range chunk {
			//fmt.Printf("%v", v)
			str += strconv.FormatUint(uint64(v), 2)
		}
		fmt.Printf("%s\n", str);
	}
}

func ExpandBitmap(data []byte) []byte{
	var tempData []byte
	for _, value := range data {
		stringy := fmt.Sprintf("%08b",value)
		for _, v := range stringy {
			switch v {
			case '1':
				tempData = append(tempData, 0xFF)
			case '0':
				tempData = append(tempData, 0x00)
			default:
				fmt.Println("unknown")
		}	
		}

	}	

	return tempData
}