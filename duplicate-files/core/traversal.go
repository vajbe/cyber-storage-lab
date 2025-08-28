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
	"sync/atomic"
	"time"
)

var (
	globalCache map[string][]string
	wg          sync.WaitGroup
	mutex       sync.Mutex
	paths       chan string
	totalFiles  atomic.Int32
)

func hashFile(dir, path string) error {

	absolutePath := dir + "/" + path

	f, err := os.Open(absolutePath)

	if err != nil {
		log.Print(err)
		return err
	}
	defer f.Close()

	h := sha256.New()

	if _, err := io.Copy(h, f); err != nil {
		log.Print(err)
		return err
	}

	hash := hex.EncodeToString(h.Sum(nil))
	mutex.Lock()
	globalCache[hash] = append(globalCache[hash], absolutePath)
	mutex.Unlock()
	totalFiles.Add(int32(1))
	return nil
}

func visitEachFile(dir string) func(path string, d fs.DirEntry, err error) error {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Print(err)
			return err
		}

		if !d.IsDir() {
			paths <- path
		}

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
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {

		for range ticker.C {
			func() {
				fmt.Println("Num GO Routines: ", runtime.NumGoroutine())
			}()
		}
	}()

	paths = make(chan string, 100)

	workers := runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				hashFile(dir, path)
			}
		}()
	}
	fs.WalkDir(fileSystem, ".", visitEachFile(dir))
	close(paths)
	wg.Wait()
	fmt.Printf("\nEnding at: %v\t Total Files Processed: %d\n", time.Since(startTime), totalFiles.Load())

	for _, value := range globalCache {
		if len(value) > 1 {
			/* fmt.Print(value) */
		}
	}

}
