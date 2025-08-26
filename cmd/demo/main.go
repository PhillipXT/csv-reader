package main

import (
	"fmt"
	"os"

	"github.com/PhillipXT/csv-reader/cmd/parser"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error checking working directory:", err)
	}
	fmt.Printf("Working directory: %s\n", dir)
	parser.LoadCSV("./assets/test.csv")
}
