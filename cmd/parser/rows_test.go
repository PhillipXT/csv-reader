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
	log.Printf("Data row: [%s]\n", data[:row1])
	assert.Equal(t, 23, row1)
	row2 := getRow(data[row1+1:])
	log.Printf("Data row: [%s]\n", data[row1+1:row1+1+row2])
	assert.Equal(t, 9, row2)

}
