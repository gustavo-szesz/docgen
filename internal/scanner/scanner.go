package scanner

import (
	"fmt"
	"log"
	"os"
)

func DirScan(dirpath string) {
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
}
