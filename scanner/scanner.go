package scanner

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func DirScan(dirpath string) {
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())

		info, err := os.Stat(dirpath)
		if err != nil {
			panic(err)
		}
		fmt.Println("Diretorio?", info.IsDir())
		if info.IsDir() {
			go_deep, err := os.Stat(dirpath)
			if err != nil {
				panic(err)
			}
			var name = fmt.Println(go_deep.Name())

			if name {
				if strings.Contains(name) {

				}
			}

			fmt.Println(go_deep.Name())

		}

	}

}
