package scanner

import (
	"fmt"
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
