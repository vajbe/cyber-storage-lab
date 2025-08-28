package core

import (
	"fmt"
	"os"
)

/* type FileChunk struct {
	path      string
	offset    int
	chunkSize int
} */

var DEFAULT_CHUNK_SIZE int = 5

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

	numWorkers := fileSize / int64(DEFAULT_CHUNK_SIZE)

	fmt.Println("Workers: ", numWorkers, " File Size: ", fileSize)

	for i := 1; i <= int(numWorkers); i++ {
		nextChunk := i * DEFAULT_CHUNK_SIZE
		buffer := make([]byte, DEFAULT_CHUNK_SIZE)
		f.ReadAt(buffer, int64(DEFAULT_CHUNK_SIZE))
		fmt.Println(string(buffer))
		fmt.Printf("\nFile: %s Chunk Size: %d Processed At: %d\n", filePath, DEFAULT_CHUNK_SIZE, nextChunk)
	}
	fmt.Println()
}
