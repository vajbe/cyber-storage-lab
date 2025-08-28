package core

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) {

	f, err := os.Open(filePath)

	if err != nil {
		fmt.Print(err)
		return
	}

	defer f.Close()
	b := make([]byte, 5000)
	f.Read(b)

	fmt.Println("File contents: ", string(b))
	fmt.Println()

}
