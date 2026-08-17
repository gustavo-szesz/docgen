# Architecture

This document describes how `docgen` is structured internally and how data flows through it.

## Goal

`docgen` parses Go source code and reconstructs a hierarchy of functions — including their signatures and the calls they make to one another — so a codebase can be documented and understood without reading it line by line.

## Packages

```
docgen/
├── analyzer/     # Parses a single .go file into structured data
├── scanner/      # Discovers .go files across a directory tree
└── cmd/docgen/   # CLI entry point — wires scanner + analyzer together
```

### `analyzer`

Responsible for everything related to **one file at a time**. It never touches the filesystem beyond reading the file it's told to parse, and it knows nothing about directories or multiple files.

**Key types**

```go
type FunctionInfo struct {
    Name    string
    Reciver string
    Params  []string
    Returns []string
    Calls   []string
}
```

**Key functions**

| Function | Responsibility |
|---|---|
| `AnalyzeFile(path string) (*ast.File, error)` | Reads a file and parses it into a Go AST |
| `GetPackageName(*ast.File) string` | Extracts the package declaration |
| `GetImports(*ast.File) []string` | Extracts the file's import paths |
| `GetFunctions(*ast.File) []FunctionInfo` | Extracts every function/method declaration, with its signature |
| `GetFunctionCalls(*ast.FuncDecl) []string` | Walks a function body and lists the functions it calls |
| `FormatType(ast.Expr) string` | Converts an AST type expression (`*ast.Ident`, `*ast.StarExpr`, `*ast.ArrayType`, ...) into a readable string |

**Data flow inside `analyzer`**

```
file path
    ↓ AnalyzeFile (os.ReadFile + go/parser)
*ast.File
    ↓ GetPackageName / GetImports / GetFunctions
package name, imports, []FunctionInfo
    ↓ (per function) GetFunctionCalls
FunctionInfo.Calls filled in
```

### `scanner`

Responsible for **finding files**, never for parsing them. It walks a directory tree and returns a list of `.go` file paths, filtering out things that shouldn't be analyzed (hidden folders, `testdata/`, `_test.go` files).

| Function | Responsibility |
|---|---|
| `FindGoFiles(root string) ([]string, error)` | Recursively lists `.go` source files under `root` |

**Data flow inside `scanner`**

```
root directory
    ↓ filepath.WalkDir
skip hidden dirs / testdata / _test.go
    ↓
[]string of .go file paths
```

### `cmd/docgen`

The CLI entry point. It's the only package that's allowed to know about *both* `scanner` and `analyzer` — it connects them: `scanner` finds the files, and each file is handed to `analyzer` to be parsed.

```
scanner.FindGoFiles(root)
    ↓
for each path:
    analyzer.AnalyzeFile(path)
    analyzer.GetFunctions(astFile)
    ↓
aggregated results across the whole project
```

## Design principles

- **Single responsibility per package.** `analyzer` only understands one file; `scanner` only understands the filesystem; `cmd/docgen` is the only place that combines them. This keeps each package testable in isolation.
- **No side effects in library code.** Parsing/scanning functions return `(result, error)` instead of calling `log.Fatal` or printing — only `cmd/docgen` (or tests) decide what to do with a failure.
- **AST-based, not text-based.** Source code is analyzed through `go/ast`, not regex or string matching, so the tool stays correct as Go syntax edge cases come up.
- **Incremental development.** Each feature (package name → imports → function signatures → function calls → directory scanning) was built and tested as its own step, with fixtures under `testdata/` per package.

## Current limitations

- `analyzer` processes one file per call; multi-file/package-wide analysis is assembled by the caller (`cmd/docgen`), not by `analyzer` itself.
- Function calls are name-based only — `docgen` does not yet resolve calls across packages or verify that a called name actually corresponds to a real function.
- Test files (`_test.go`) are excluded from scanning by design, for now.

## Roadmap

- Build a real call hierarchy/tree from the flat `Calls` lists.
- Generate human-readable documentation output (Markdown/HTML) from the analyzed data.
- Wire `scanner` + `analyzer` together fully inside `cmd/docgen`.
