package scanner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var listDocs []string

func DirScan(dirpath string) []string {
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

			var aux string = filepath.Join(dirpath, e.Name())
			DirScan(aux)
			fmt.Print(dirpath + "/" + e.Name()) // TODO: Improve for Windows Users
		} else {
			if strings.HasSuffix(e.Name(), ".go") {
				listDocs = append(listDocs, e.Name())

				fmt.Println("Found:", e.Name())
			}
		}

	}
	return listDocs

}
