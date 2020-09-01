package search

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
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
	fmt.Println("\t\tSearch results")
	fmt.Println("--------------------------------------------------------------")
	for _, fileName := range filesAndFolders {
		fmt.Printf("%v\n", fileName)
	}
}

func getMatchingFilesRecursive(filename string, directory string, skipDirectory string, match string) ([]string, []int64, error) {
	filesAndFolders := []string{}
	// index 1 is number of files searched index 2 is number of file that match
	fileNumbers := []int64{0, 0}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, nil, err
	}

	for _, file := range files {
		fileNumbers[0]++
		name := file.Name()
		nameWithPath := directory + "\\" + name

		if match == "exact" {
			if name == filename {
				fileNumbers[1]++
				filesAndFolders = append(filesAndFolders, nameWithPath)
			}
		} else {
			if strings.Contains(name, filename) {
				fileNumbers[1]++
				filesAndFolders = append(filesAndFolders, nameWithPath)
			}
		}

		if file.IsDir() && name != skipDirectory {
			dir := directory + "\\" + name
			names, fileSearched, err := getMatchingFilesRecursive(filename, dir, skipDirectory, match)
			fileNumbers[0] = fileNumbers[0] + fileSearched[0]
			fileNumbers[1] = fileNumbers[1] + fileSearched[1]
			if err != nil {
				return nil, nil, err
			}

			filesAndFolders = append(filesAndFolders, names...)
		}
	}

	return filesAndFolders, fileNumbers, nil

}

func FindFileByNameRecursive(filename string, directory string, skipDirectory string, match string) {
	t := time.Now()
	filesAndFolders := []string{}

	filesAndFolders, fileNumbers, err := getMatchingFilesRecursive(filename, directory, skipDirectory, match)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\t\tSearch results")
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Time Taken to search: %v, Total Files Scanned: %v , Total Files Found: %v \n", t.Sub(time.Now()), fileNumbers[0], fileNumbers[1])
	for _, fileName := range filesAndFolders {
		fmt.Printf("%v\n", fileName)
	}
	fmt.Println("--------------------------------------------------------------")
}
