package scanner

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func DirScan(dirpath string) {
	var listDocs []string
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())

		if e.IsDir() {
			if strings.Contains(e.Name(), ".go") {
				listDocs = append(listDocs, e.Name())
			}
		}

	}

}
