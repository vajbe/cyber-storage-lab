package main

import (
	"find-duplicates-files/core"
)

func main() {
	/* dir := flag.String("d", "dir", "directory path")
	flag.Parse()
	if *dir == "dir" {
		log.Fatal("Directory path is missing.")
	} */
	// core.TraverseDirectory(*dir)

	core.ReadFile("/home/vivek/codes/cyber-storage-lab/find-duplicate-files/go.mod")
}
