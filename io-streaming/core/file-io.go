package core

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

/* type FileChunk struct {
	path      string
	offset    int
	chunkSize int
} */

//Processing around 10GB of file

var (
	DEFAULT_CHUNK_SIZE = 4 * 1024 * 1024
	wg                 sync.WaitGroup
)

func ReadByChunks(offset int64, f *os.File, chunkSize int) {
	buffer := make([]byte, chunkSize)
	_, err := f.ReadAt(buffer, int64(offset))
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Read error:", err)
	}

	//Wasting CPU stimulating various work

	/* h := sha256.Sum256(buffer[:n])
	for i := 0; i < 100000; i++ {
		big.NewInt(0).Exp(big.NewInt(2), big.NewInt(20), nil)
	}
	_ = h */
	time.Sleep(1 * time.Second)
}

func ReadFile(filePath string) {

	f, err := os.Open(filePath)

	if err != nil {
		fmt.Print(err)
		return
	}

	defer f.Close()

	stat, err := os.Stat(filePath)
	if err != nil {
		fmt.Print(err)
		return
	}

	fileSize := stat.Size()

	numChunks := (fileSize + int64(DEFAULT_CHUNK_SIZE) - 1) / int64(DEFAULT_CHUNK_SIZE)

	fmt.Println("Chunks:", numChunks, "File Size:", fileSize, "Bytes")

	nw := time.Now()

	numCpu := runtime.NumCPU()
	chunks := make(chan int64, numCpu*8)
	for i := 0; i < numCpu; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for offset := range chunks {
				ReadByChunks(offset, f, DEFAULT_CHUNK_SIZE)
			}
		}()
	}

	for i := int64(0); i < numChunks; i++ {
		offset := i * int64(DEFAULT_CHUNK_SIZE)
		chunks <- offset
	}

	close(chunks)
	wg.Wait()

	fmt.Print("Time taken: ", time.Since(nw))
}
