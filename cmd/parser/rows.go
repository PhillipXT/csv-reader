package parser

import (
	"fmt"
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

	for {
		if i >= len(data) {
			break
		}
		if data[i] == dquote {
			if colPos == 0 {
				quoteMode = true
			} else if !quoteMode {
				log.Fatal("Found double quotes in a non-quoted column.")
			} else {
				if data[i+1] == dquote {
					i += 1
				} else if data[i+1] == delim {
					quoteMode = false
					colPos = -1
					i += 1
				} else if data[i+1] == byte('\n') {
					break
				} else {
					log.Fatal("Expected new column, found extra characters instead.")
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

	fmt.Println("Line:", string(data[:i]))

	return i
}
