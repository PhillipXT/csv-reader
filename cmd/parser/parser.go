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
	eof := false

	for {
		if pos >= len(buf) {
			new_buf := make([]byte, len(buf)*2)
			copy(new_buf, buf)
			buf = new_buf
		}

		//fmt.Printf("Buffer size: %d\n", len(buf))

		bytesRead, err := file.Read(buf[pos:])
		if err == io.EOF {
			fmt.Println("End of file reached.")
			eof = true
		}
		if err != nil {
			log.Fatal("Error reading file:", err)
		}

		pos += bytesRead

		bytesUsed, err := processData(buf[:pos], eof)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		if bytesUsed > 0 {
			copy(buf, buf[bytesUsed:])
			pos -= bytesUsed
		}

		if eof {
			break
		}
	}
}

func processData(data []byte, eof bool) (int, error) {
	totalBytesUsed := 0

	if len(data) == 0 {
		return 0, nil
	}

	for {
		bytesUsed := getRow(data[totalBytesUsed:], eof)
		if bytesUsed == 0 {
			break
		}

		totalBytesUsed += bytesUsed
		fmt.Printf("bytesUsed: %d\n", bytesUsed)
		row := parseColumns(data[:bytesUsed])
		row.Print()
	}

	return totalBytesUsed, nil
}
