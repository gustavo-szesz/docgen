package scanner

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DirScan(dirpath string) []string {
	var listDocs []string
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
		if e.IsDir() {

			if e.Name() == ".git" {
				continue
			}

			newPath := filepath.Join(dirpath, e.Name())
			findFiles := DirScan(newPath)

			listDocs = append(listDocs, findFiles...)

		} else {
			if strings.HasSuffix(e.Name(), ".go") {

				completePath := filepath.Join(
					dirpath,
					e.Name(),
				)
				listDocs = append(listDocs, completePath)

				fmt.Println("Found:", completePath)
			}
		}

	}
	return listDocs

}

func FindGoFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") || name == "testdata" {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}
