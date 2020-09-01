package search

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// FileInfoByName Gets Filedetails by full file path
func FileInfoByName(filename string) {
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	name := fileInfo.Name()
	size := fileInfo.Size()
	mode := fileInfo.Mode()
	modifiedOn := fileInfo.ModTime()
	isDir := fileInfo.IsDir()
	// sys := fileInfo.Sys()

	// Print file details
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\t\tFile Details")
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Name: \t\t%v\n", name)
	fmt.Printf("Size: \t\t%v\n", size)
	fmt.Printf("Permissions: \t%v\n", mode)
	fmt.Printf("Modified On: \t%v\n", modifiedOn)
	fmt.Printf("Is Folder: \t%v\n", isDir)
	fmt.Println("--------------------------------------------------------------")
	// fmt.Println("Other Details: %v", sys)
}

func FindFileByName(filename string, directory string) {
	filesAndFolders := []string{}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		name := file.Name()
		if strings.Contains(name, filename) {
			filesAndFolders = append(filesAndFolders, name)
		}
	}

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\t\tFile Matches")
	fmt.Println("--------------------------------------------------------------")
	for _, fileName := range filesAndFolders {
		fmt.Printf("%v\n", fileName)
	}
}
