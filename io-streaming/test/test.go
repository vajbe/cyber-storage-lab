package test

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "

// randomString generates a random string of given length
func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateFile() {
	fileName := "large-file.txt"
	targetSize := int64(5 * 1024 * 1024 * 1024) // 10 GB
	chunkSize := 1024 * 1024                    // 1 MB per write

	rand.Seed(time.Now().UnixNano())

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := bufio.NewWriterSize(f, chunkSize)

	var written int64
	start := time.Now()

	for written < targetSize {
		line := randomString(100) + "\n"
		for writer.Buffered() < chunkSize {
			_, _ = writer.WriteString(line)
			written += int64(len(line))
			if written >= targetSize {
				break
			}
		}
		_ = writer.Flush()
	}

	fmt.Printf("Generated file: %s (%.2f GB) in %v\n", fileName, float64(written)/(1024*1024*1024), time.Since(start))
}
