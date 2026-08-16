package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type FunctionInfo struct {
	Name    string
	Reciver string
	Params  []string
	Returns []string
	Calls   []string
}

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

	// pkg_name := GetPackageName(astFile)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("%s\n", pkg_name)
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

func GetFunctions(funcDecl *ast.FuncDecl) []string {
	//listFunctions := []FunctionInfo{}
	var calls []string

	if funcDecl.Body == nil {
		return calls
	}

	ast.Inspect(funcDecl.Body, func(node ast.Node) bool {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		name := extracCallName(callExpr.Fun)
		if name != "" {
			calls = append(calls, name)
		}
		return true
	})

	return calls
}

func extracCallName(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.SelectorExpr:
		return FormatType(v.X) + "." + v.Sel.Name
	default:
		return ""
	}
}
func FormatType(expre ast.Expr) string {
	switch v := expre.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.StarExpr:
		return "*" + FormatType(v.X)
	case *ast.ArrayType:
		return "[]" + FormatType(v.Elt)

	default:
		return "unknown"
	}
}
