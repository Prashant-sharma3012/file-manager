package search

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync/atomic"
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

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("\t\tSearch results")
	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Printf("Time Taken to search: %v, Total Files Scanned: %v , Total Files Found: %v \n", time.Now().Sub(t), fileNumbers[0], fileNumbers[1])
	fmt.Println("----------------------------------------------------------------------------------------")
	for _, fileName := range filesAndFolders {
		fmt.Printf("%v\n", fileName)
	}
	fmt.Println("----------------------------------------------------------------------------------------")
}

type dirScanDetails struct {
	filesAndFolders []string
	fileNumbers     []int64
	err             error
}

func getMatchingFilesConcurrent(filename string, directory string, skipDirectory string, match string, c chan dirScanDetails, running *int64) {

	details := dirScanDetails{
		filesAndFolders: []string{},
		fileNumbers:     []int64{0, 0},
		err:             nil,
	}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		atomic.AddInt64(running, -1)
		c <- details
		return
	}

	for _, file := range files {
		details.fileNumbers[0]++
		name := file.Name()
		nameWithPath := directory + "\\" + name

		if match == "exact" {
			if name == filename {
				details.fileNumbers[1]++
				details.filesAndFolders = append(details.filesAndFolders, nameWithPath)
			}
		} else {
			if strings.Contains(name, filename) {
				details.fileNumbers[1]++
				details.filesAndFolders = append(details.filesAndFolders, nameWithPath)
			}
		}

		if file.IsDir() && name != skipDirectory {
			dir := directory + "\\" + name
			atomic.AddInt64(running, 1)
			go getMatchingFilesConcurrent(filename, dir, skipDirectory, match, c, running)
		}
	}

	atomic.AddInt64(running, -1)
	c <- details
	return
}

func FindFileByNameConcurrent(filename string, directory string, skipDirectory string, match string) {
	t := time.Now()
	c := make(chan dirScanDetails, 100)
	var running int64

	details := dirScanDetails{
		filesAndFolders: []string{},
		fileNumbers:     []int64{0, 0},
		err:             nil,
	}

	atomic.AddInt64(&running, 1)
	go getMatchingFilesConcurrent(filename, directory, skipDirectory, match, c, &running)

	for v := range c {
		if v.err != nil {
			fmt.Println(v.err)
			break
		} else {
			details.filesAndFolders = append(details.filesAndFolders, v.filesAndFolders...)
			details.fileNumbers[0] = details.fileNumbers[0] + v.fileNumbers[0]
			details.fileNumbers[1] = details.fileNumbers[1] + v.fileNumbers[1]
		}
		if running == 0 {
			break
		}
	}

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("\t\tSearch results")
	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Printf("Time Taken to search: %v, Total Files Scanned: %v , Total Files Found: %v \n", time.Now().Sub(t), details.fileNumbers[0], details.fileNumbers[1])
	fmt.Println("----------------------------------------------------------------------------------------")
	for _, fileName := range details.filesAndFolders {
		fmt.Printf("%v\n", fileName)
	}
	fmt.Println("----------------------------------------------------------------------------------------")
}
