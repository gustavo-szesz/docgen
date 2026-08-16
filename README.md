# docgen

docgen is a command-line tool written in Go that analyzes Go source code and organizes its functions into a hierarchy, making it easier to read, document, and navigate large codebases.

Instead of scrolling through hundreds of lines of source, docgen inspects your code with Go's AST (go/ast) and produces a structured view of what each file contains and how its functions relate to one another.
```
⚠️ Status: This project is under active development. Core parsing features are being built incrementally, issue by issue.
```
What it does (so far)

Given a .go source file, docgen can extract:

Package name — which package the file belongs to.
Imports — the list of packages it depends on.
Function signatures — name, receiver (for methods), parameters, and return types.
Function calls — which other functions each function calls, laying the groundwork for a full call hierarchy.
Planned features
Recursively scan an entire directory (not just a single file) for .go sources.
Build a full call hierarchy across functions and files.
Generate readable documentation output (e.g. Markdown or HTML) from the analyzed structure.
Project structure
```
docgen/
├── analyzer/     # Parses a single .go file and extracts package, imports, functions, and calls
├── scanner/      # (in progress) Discovers .go files across a project directory
├── cmd/docgen/   # CLI entry point
└── go.mod
How it works (high level)
Go source file
      ↓
os.ReadFile
      ↓
go/parser → *ast.File
      ↓
analyzer.GetPackageName / GetImports / GetFunctions
      ↓
structured function data (name, params, returns, calls)
Development approach
```
This project is built as a learning exercise, one GitHub issue at a time. Each issue introduces one small, focused piece of functionality (e.g. "extract package name", "extract function signatures", "detect function calls"), with tests added alongside each feature under testdata/.

Running tests
```
bash
go test ./... -v
```
License

Not yet defined.
