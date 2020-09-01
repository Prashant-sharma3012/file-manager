package main

import (
	"flag"

	"github.com/Prashant-sharma3012/file-manager/search"
)

func main() {

	command := flag.String("op", " ", "Top Level Command")
	fileName := flag.String("fname", " ", "File name whose details are needed")
	dir := flag.String("dir", " ", "File name whose details are needed")
	mode := flag.String("r", "false", "File name whose details are needed")
	skipDir := flag.String("skipDir", " ", "Folders not to be searched")
	match := flag.String("match", " ", "Folders not to be searched")

	flag.Parse()

	switch true {
	case *command == "details":
		search.FileInfoByName(*fileName)
	case *command == "search":
		if *mode == "true" {
			// search.FindFileByNameRecursive(*fileName, *dir, *skipDir, *match)
			search.FindFileByNameConcurrent(*fileName, *dir, *skipDir, *match)
		} else {
			search.FindFileByName(*fileName, *dir)
		}
	}
}
