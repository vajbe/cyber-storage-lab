package core

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func TraverseDirectory(dir string) {
	fileSystem := os.DirFS(dir)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(path)
		return nil
	})
}
