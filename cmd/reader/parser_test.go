package loader

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestRowParser(t *testing.T) {

	// Test: Valid single row
	log.Println("Testing valid single row...")
	data := []byte("One,Two,Three,Four,Five")
	row := parseLine(data)
	row.Print()
	assert.Equal(t, 5, row.columnCount)
	assert.Equal(t, []string{"One", "Two", "Three", "Four", "Five"}, row.columns)

	// Test: Valid field with a comma
	log.Println("Testing valid field with a comma...")
	data = []byte("One,Two,Three,Four,\"Five, Six\"")
	row = parseLine(data)
	row.Print()
	assert.Equal(t, 5, row.columnCount)
	assert.Equal(t, []string{"One", "Two", "Three", "Four", "Five, Six"}, row.columns)

	// Test: Valid field with comma and double quotes at beginning
	log.Println("Testing valid field with comma and double quotes...")
	data = []byte("\"One, \"\"Two\"\"\",Three,Four,Five,Six")
	row = parseLine(data)
	row.Print()
	assert.Equal(t, 5, row.columnCount)
	assert.Equal(t, []string{"One, \"Two\"", "Three", "Four", "Five", "Six"}, row.columns)

	// Test: Valid field with comma and double quotes at beginning and end
	log.Println("Testing valid fields with commas and double quotes...")
	data = []byte("\"\"\"One, Two\"\", Three\",Four,\"Five, \"\"Six\"\"\"")
	row = parseLine(data)
	row.Print()
	assert.Equal(t, 3, row.columnCount)
	assert.Equal(t, []string{"\"One, Two\", Three", "Four", "Five, \"Six\""}, row.columns)

}
