package analyzer

import (
	"testing"
)

func TestGetPackageName(t *testing.T) {
	astFile, err := AnalyzeFile("analyzer/testdata/sample.go")
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
