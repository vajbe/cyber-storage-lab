package main

import (
	"find-duplicates-files/core"
	"flag"
	"log"
)

func main() {
	dir := flag.String("d", "dir", "directory path")
	flag.Parse()
	if *dir == "dir" {
		log.Fatal("Directory path is missing.")
	}
	core.TraverseDirectory(*dir)
}
