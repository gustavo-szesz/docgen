package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gustavo-szesz/docgen/scanner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing arguments.")
		fmt.Println("Usage: go run main.go [dirpath]")
		return
	}
	aux := os.Args[1]
	mapped_directorys := scanner.DirScan(aux)
	fmt.Println(mapped_directorys)

	for _, file := range mapped_directorys {
		content, err := os.ReadFile(file)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(content))
	}

}
