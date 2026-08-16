package analyzer

import (
	"testing"
)

func TestGetPackageName(t *testing.T) {
	astFile, err := AnalyzeFile("testdata/sample.go")
	if err != nil {
		t.Fatalf("Not suposed to have error, %v.", err)
	}
	// Act
	result := GetPackageName(astFile)

	// Assert
	waited_result := "sample"
	if result != waited_result {
		t.Errorf("Waited result %q, Recive %q", waited_result, result)
	}
}

func TestAnalyzeFile_NotFoundFile(t *testing.T) {
	// Act
	_, err := AnalyzeFile("testdata/no_file.go")

	if err == nil {
		t.Error("Not Expected a error, err as nil")
	}
}

func TestGetImports(t *testing.T) {
	astFile, err := AnalyzeFile("testdata/with_imports.go")
	if err != nil {
		t.Fatalf("Not expected error: %v", err)
	}
	result := GetImports(astFile)

	waited_result := []string{`"fmt"`, `"ok"`}

	if len(result) != len(waited_result) {
		t.Fatalf("Waited results %d import, came %d", len(waited_result), len(result))
	}

	for i := 0; i > len(waited_result); i++ {
		if result[i] != waited_result[i] {
			t.Errorf("import[%d]: waited %q, result %q", i, waited_result[i], result[i])
		}
	}

}
