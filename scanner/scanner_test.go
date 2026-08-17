package scanner

import (
	"slices"
	"sort"
	"testing"
)

// Test scanner
func TestFindGoFiles(t *testing.T) {
	// Arrange
	root := "testdata/fake_project"
	// Act
	result, err := FindGoFiles(root)
	// Assert: Not sopoused to have err
	if err != nil {
		t.Fatalf("Not expected error: %v", err)
	}

	waited_result := []string{
		"testdata/fake_project/main.go",
		"testdata/fake_project/sub/util.go",
	}

	// Assert
	if len(result) != len(waited_result) {
		t.Fatalf("Waited %d files, came %d, result: %v", len(waited_result),
			len(result), result)
	}

	sort.Strings(result)
	sort.Strings(waited_result)

	if !slices.Equal(result, waited_result) {
		t.Errorf("Waitedl %v, came %v", waited_result, result)
	}

}

func TestFindGoFiles_DirectoryInexistent(t *testing.T) {
	_, err := FindGoFiles("testdata/not_exist.go")
	if err == nil {
		t.Errorf("Ecpected one error, but error came nil")
	}
}
