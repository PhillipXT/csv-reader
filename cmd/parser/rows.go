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

func getRow(data []byte, eof bool) int {

	quoteMode := false
	colPos := 0
	i := 0
	foundRow := false
	hasLinefeed := false

	if len(data) == 0 {
		return 0
	}

	log.Printf("Processing data: %s (%d)", data, len(data))
	for {
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
					hasLinefeed = true
					foundRow = true
					break
				} else if i >= len(data)-1 && eof == true {
					foundRow = true
					break
				} else if i >= len(data)-1 {
					break
				} else {
					log.Fatalf("Expected new column, found extra characters instead (%d = %s/%s)", i, string(data[i]), string(data[i+1]))
				}
			}
		} else if !quoteMode && data[i] == delim {
			quoteMode = false
			colPos = -1
		} else if !quoteMode && data[i] == byte('\n') {
			hasLinefeed = true
			foundRow = true
			break
		}
		if i >= len(data)-1 && !foundRow && eof && quoteMode {
			log.Fatalf("Malformed column - no end quote found")
		} else if i >= len(data)-1 {
			if eof == true {
				foundRow = true
			}
			break
		}
		i += 1
		colPos += 1
	}

	if foundRow {
		if eof && !hasLinefeed {
			return i
		}
		return i + 1
	}

	return 0
}
