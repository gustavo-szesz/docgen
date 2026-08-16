package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func AnalyzeFile(file string) (*ast.File, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	fileset := token.NewFileSet()

	astFile, err := parser.ParseFile(
		fileset,
		file,
		content,
		parser.ParseComments,
	)
	if err != nil {
		return nil, err
	}
	return astFile, nil
}

func GetPackageName(file *ast.File) string {
	return file.Name.Name
}

func GetImports(astFile *ast.File) []string {
	var listImports []string

	for _, e := range astFile.Imports {
		if e.Path != nil {
			pathValue := e.Path.Value

			listImports = append(listImports, pathValue)
		}
	}
	return listImports
}

// func GetFunctions(astFile *ast.File) {
// 	var listFunctions []string

// 	for _, e := range astFile.Decls {
// 		if e.Name()
// 	}
// }
