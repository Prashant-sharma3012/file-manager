package main

import (
	"flag"

	"github.com/Prashant-sharma3012/file-manager/search"
)

func main() {

	command := flag.String("op", " ", "Top Level Command")
	fileName := flag.String("fname", " ", "File name whose details are needed")
	dir := flag.String("dir", " ", "File name whose details are needed")

	flag.Parse()

	switch true {
	case *command == "details":
		search.FileInfoByName(*fileName)
	case *command == "search":
		search.FindFileByName(*fileName, *dir)
	}
}
