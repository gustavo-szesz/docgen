package main

import (
	"fmt"
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
	scanner.DirScan(aux)

	// conteudo, err := os.ReadFile(aux)
	// if err != nil {
	// 	log.Fatal(err)

	// }
	// fmt.Println(string(conteudo))
}
