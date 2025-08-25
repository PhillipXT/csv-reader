package loader

import (
	"bytes"
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
				readLine(buf[:bytesRead])
				fmt.Println("End of file reached.")
				break
			}
			log.Fatal("Error reading file:", err)
		}

		pos += bytesRead

		bytesUsed := readLine(buf[:pos])

		copy(buf, buf[bytesUsed:])

		pos -= bytesUsed
	}
}

func readLine(line []byte) int {

	linefeed := bytes.Index(line, []byte("\n"))
	if linefeed < 0 {
		return 0
	}

	fmt.Println("Line:", string(line[:linefeed]))

	return linefeed + 1
}
