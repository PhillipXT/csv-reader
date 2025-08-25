package main

import (
	"fmt"
	"os"

	loader "github.com/PhillipXT/csv-reader/cmd/reader"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error checking working directory:", err)
	}
	fmt.Printf("Working directory: %s\n", dir)
	loader.LoadCSV("./assets/test.csv")
}
