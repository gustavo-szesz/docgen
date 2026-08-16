package main

import (
	"fmt"
	"os"

	"github.com/gustavo-szesz/docgen/analyzer"
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

	analyzer.AnalyzeFile(aux)
}

// reading the doc, TODO: SRP
// for _, file := range mapped_directorys {
// 	content, err := os.ReadFile(file)
// 	if err != nil {
// 		break
// 	}

// 	fileset := token.NewFileSet()

// 	ast, err := parser.ParseFile(
// 		fileset,
// 		file,
// 		content,
// 		parser.ParseComments,
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("==========================")
// 	fmt.Println("File: " + file)
// 	fmt.Println("==========================")
// 	fmt.Println("Packege:", ast.Name)
