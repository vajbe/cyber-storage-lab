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
	"sync"
	"time"
)

var (
	globalCache map[string][]string
	wg          sync.WaitGroup
	mutex       sync.Mutex
)

func visitEachFile(dir string, wg *sync.WaitGroup) func(path string, d fs.DirEntry, err error) error {
	return func(path string, d fs.DirEntry, err error) error {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var m runtime.MemStats
			if err != nil {
				log.Fatal(err)
			}

			if !d.IsDir() {
				runtime.ReadMemStats(&m)
				fmt.Println("Started reading: \t", d.Name(), " \tat ", time.Now(), "\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
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
				mutex.Lock()
				globalCache[hash] = append(globalCache[hash], absolutePath)
				mutex.Unlock()
				runtime.ReadMemStats(&m)
				fmt.Println("Finished reading: \t", d.Name(), " \tat ", time.Now(), "\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
			}
		}()
		return nil
	}
}

func TraverseDirectory(dir string) {

	globalCache = make(map[string][]string)
	fileSystem := os.DirFS(dir)
	startTime := time.Now()
	wg = sync.WaitGroup{}
	mutex = sync.Mutex{}

	fmt.Println("Starting process: ", startTime)
	fs.WalkDir(fileSystem, ".", visitEachFile(dir, &wg))
	wg.Wait()
	fmt.Println("Ending at: ", time.Since(startTime))
	for _, value := range globalCache {
		if len(value) > 1 {
			fmt.Print(value)
		}
	}

}
