package parser

import (
	"log"
)

type readerState int

const (
	stateInsideRow  readerState = 1
	stateInnerQuote readerState = 2
)

const dquote byte = byte('"')
const delim = byte(',')

func getRow(data []byte) int {

	quoteMode := false
	colPos := 0
	i := 0

	log.Printf("Processing data: %s (%d)", data, len(data))
	for {
		if i >= len(data) {
			i -= 1
			log.Printf("Returning %d\n", i+1)
			break
		}
		//log.Printf("Processing data: %d = %s\n", i, string(data[i]))
		if data[i] == dquote {
			if colPos == 0 {
				quoteMode = true
			} else if !quoteMode {
				log.Fatal("Found double quotes in a non-quoted column.")
			} else {
				if i < len(data)-1 && data[i+1] == dquote {
					i += 1
				} else if i < len(data)-1 && data[i+1] == delim {
					quoteMode = false
					colPos = -1
					i += 1
				} else if i < len(data)-1 && data[i+1] == byte('\n') {
					break
				} else if i == len(data)-1 {
					break
				} else {
					log.Fatalf("Expected new column, found extra characters instead (%d = %s/%s)", i, string(data[i]), string(data[i+1]))
				}
			}
		} else if !quoteMode && data[i] == delim {
			quoteMode = false
			colPos = -1
		} else if !quoteMode && data[i] == byte('\n') {
			break
		}
		i += 1
		colPos += 1
	}

	return i + 1
}
