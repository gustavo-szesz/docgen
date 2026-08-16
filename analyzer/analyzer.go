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

func GetFunctionsCalls(astFile *ast.File) []FunctionInfo {
	listFunctions := []FunctionInfo{}

	for _, e := range astFile.Decls {
		funcDect, ok := e.(*ast.FuncDecl)
		if !ok {
			continue
		}

		info := FunctionInfo{}
		info.Name = funcDect.Name.Name

		if funcDect.Recv != nil && len(funcDect.Recv.List) > 0 {
			info.Reciver = FormatType(funcDect.Recv.List[0].Type)
		}

		if funcDect.Type.Params != nil {
			for _, field := range funcDect.Type.Params.List {
				typeText := FormatType(field.Type)
				for _, name := range field.Names {
					info.Params = append(info.Params, name.Name+" "+typeText)
				}
			}
		}
		if funcDect.Type.Results != nil {
			for _, field := range funcDect.Type.Results.List {
				info.Returns = append(info.Returns, FormatType(field.Type))
			}
		}

		listFunctions = append(listFunctions, info)

	}
	return listFunctions
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
