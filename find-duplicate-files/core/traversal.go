package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"runtime"
	"time"
)

var globalCache map[string][]string

func TraverseDirectory(dir string) {
	var m runtime.MemStats
	globalCache = make(map[string][]string)
	fileSystem := os.DirFS(dir)
	startTime := time.Now()
	fmt.Println("Starting process: ", startTime)
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !d.IsDir() {
			fmt.Println("Started reading: \t", d.Name(), " \tat ", time.Now(), "\tNumGC = ", m.NumGC)
			absolutePath := dir + "/" + path

			f, err := os.Open(absolutePath)

			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			h := sha256.New()

			if _, err := io.Copy(h, f); err != nil {
				log.Fatal(err)
			}

			hash := hex.EncodeToString(h.Sum(nil))

			if _, ok := globalCache[hash]; !ok {
				globalCache[hash] = []string{}
			}

			globalCache[hash] = append(globalCache[hash], absolutePath)
			fmt.Println("Finished reading: \t", d.Name(), " \tat ", time.Now(), "\tNumGC = ", m.NumGC)
		}
		return nil
	})

	for _, value := range globalCache {
		if len(value) > 1 {
			fmt.Print(value)
		}
	}

	fmt.Println("Ending at: ", time.Since(startTime))

}
