package parser

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestColumnParser(t *testing.T) {

	// Test: Valid single row
	log.Println("Testing valid single row...")
	data := []byte(`One,Two,Three,Four,Aa09!Øڃಚ𠜎😁`)
	row := parseColumns(data)
	row.Print()
	assert.Equal(t, 5, row.columnCount)
	assert.Equal(t, []string{`One`, `Two`, `Three`, `Four`, `Aa09!Øڃಚ𠜎😁`}, row.columns)

	// Test: Valid field with a comma
	log.Println("Testing valid field with a comma...")
	data = []byte(`One,Two,Three,Four,"Five, Six"`)
	row = parseColumns(data)
	row.Print()
	assert.Equal(t, 5, row.columnCount)
	assert.Equal(t, []string{`One`, `Two`, `Three`, `Four`, `Five, Six`}, row.columns)

	// Test: Valid field with comma and double quotes at beginning
	log.Println("Testing valid field with comma and double quotes...")
	data = []byte(`"One, ""Two""",Three,Four,Five,Six`)
	row = parseColumns(data)
	row.Print()
	assert.Equal(t, 5, row.columnCount)
	assert.Equal(t, []string{`One, "Two"`, `Three`, `Four`, `Five`, `Six`}, row.columns)

	// Test: Valid field with comma and double quotes at beginning and end
	log.Println("Testing valid fields with commas and double quotes...")
	data = []byte(`"""One, Two"", Three",Four,"Five, ""Six"""`)
	row = parseColumns(data)
	row.Print()
	assert.Equal(t, 3, row.columnCount)
	assert.Equal(t, []string{`"One, Two", Three`, `Four`, `Five, "Six"`}, row.columns)

	// Test: Valid single row with line break inside a field
	log.Println("Testing valid single row with a line break in a field...")
	data = []byte("One,Two,Three,Four,Five,\"line\nbreak\"")
	row = parseColumns(data)
	row.Print()
	assert.Equal(t, 6, row.columnCount)
	assert.Equal(t, []string{"One", "Two", "Three", "Four", "Five", "line\nbreak"}, row.columns)

}
