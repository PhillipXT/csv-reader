package parser

import (
	"fmt"
	"io"
	"log"
	"os"
)

const bufferSize = 8

func LoadCSV(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	defer file.Close()

	buf := make([]byte, bufferSize)
	pos := 0

	for {
		if pos >= len(buf) {
			new_buf := make([]byte, len(buf)*2)
			copy(new_buf, buf)
			buf = new_buf
		}

		//fmt.Printf("Buffer size: %d\n", len(buf))

		bytesRead, err := file.Read(buf[pos:])
		if err != nil {
			if err == io.EOF {
				getRow(buf[:bytesRead])
				fmt.Println("End of file reached.")
				break
			}
			log.Fatal("Error reading file:", err)
		}

		pos += bytesRead

		bytesUsed := getRow(buf[:pos])
		row := parseColumns(buf[:bytesUsed])
		row.Print()

		copy(buf, buf[bytesUsed:])

		pos -= bytesUsed
	}
}
