package parser

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRowParser(t *testing.T) {

	// Test: Valid single row
	log.Println("Testing valid single row...")
	data := []byte("One,Two,Three,Four,Five\n1,2,3,4,5")
	row1 := getRow(data)
	log.Printf("Data row: 0-%d [%s]\n", row1, data[:row1])
	assert.Equal(t, 24, row1)
	row2 := getRow(data[row1:])
	log.Printf("Data row: %d-%d [%s]\n", row1, row2, data[row1:row1+row2])
	assert.Equal(t, 9, row2)

	// Test: Valid single row
	log.Println("Testing valid single row with quotes and newline...")
	data = []byte("One,Two,Three,Four,Five\n1,2,3,4,\"\"\"5\n6\"\"\"")
	row1 = getRow(data)
	log.Printf("Data row: 0-%d [%s]\n", row1, data[:row1])
	assert.Equal(t, 24, row1)
	row2 = getRow(data[row1:])
	log.Printf("Data row: %d-%d [%s]\n", row1, row2, data[row1:row1+row2])
	assert.Equal(t, 17, row2)

}
