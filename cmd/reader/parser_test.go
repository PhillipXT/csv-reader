package loader

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestRowParser(t *testing.T) {

	// Test: Valid single header
	log.Println("Testing valid single header...")
	data := []byte("One,Two,Three,Four,Five")
	row := parseLine(data)
	row.Print()
	assert.Equal(t, 5, row.columnCount)
	assert.Equal(t, []string{"One", "Two", "Three", "Four", "Five"}, row.columns)

}
