package main

import (
	"flag"

	"github.com/Prashant-sharma3012/file-manager/archive"
	"github.com/Prashant-sharma3012/file-manager/search"
)

func main() {

	command := flag.String("op", " ", "Top Level Command")
	fileName := flag.String("fname", " ", "File name whose details are needed")
	dir := flag.String("dir", " ", "File name whose details are needed")
	mode := flag.String("r", "false", "File name whose details are needed")
	skipDir := flag.String("skipDir", " ", "Folders not to be searched")
	match := flag.String("match", " ", "Folders not to be searched")
	multiThread := flag.String("multiThread", " ", "Folders not to be searched")
	filepath := flag.String("fpath", " ", "File name whose details are needed")

	flag.Parse()

	switch true {
	case *command == "details":
		search.FileInfoByName(*fileName)
	case *command == "search":
		if *mode == "true" {
			if *multiThread == "true" {
				search.FindFileByNameConcurrent(*fileName, *dir, *skipDir, *match)
			} else {
				search.FindFileByNameRecursive(*fileName, *dir, *skipDir, *match)
			}
		} else {
			search.FindFileByName(*fileName, *dir)
		}
	case *command == "zip":
		archive.ZipFile(*filepath, *dir)
	}
}
