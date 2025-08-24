package core

import (
	"crypto/sha256"
	"fmt"
	"io/fs"
	"log"
	"os"
)

var globalCache map[string][]string

func TraverseDirectory(dir string) {
	globalCache = make(map[string][]string)
	fileSystem := os.DirFS(dir)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !d.IsDir() {
			absolutePath := dir + "/" + path
			content, err := os.ReadFile(absolutePath)
			if err != nil {
				log.Fatal(err)
			}

			hashBytes := sha256.Sum256(content)
			hash := string(hashBytes[:])

			if _, ok := globalCache[hash]; !ok {
				globalCache[hash] = []string{}
			}

			globalCache[hash] = append(globalCache[hash], absolutePath)

		}
		return nil
	})

	for _, value := range globalCache {
		if len(value) > 1 {
			fmt.Print(value)
		}
	}

}
