package archive

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync/atomic"

	"github.com/Prashant-sharma3012/file-manager/utils"
)

func readFileDataAndInfo(filePath string) ([]byte, os.FileInfo, error) {
	// read file bytes
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, nil, err
	}

	return fileData, fileInfo, nil
}

func getCompressor(dest string) (*os.File, *zip.Writer, error) {
	fd, err := os.Create(dest)
	if err != nil {
		return nil, nil, err
	}

	compressor := zip.NewWriter(fd)
	return fd, compressor, nil
}

func makeFileHeader(info os.FileInfo) (*zip.FileHeader, error) {
	fileHeader, err := zip.FileInfoHeader(info)
	if err != nil {
		return nil, err
	}

	fileHeader.Method = zip.Deflate

	return fileHeader, nil
}

func compress(fileHeader *zip.FileHeader, fileData []byte, compressor *zip.Writer) (int, error) {
	compressed, err := compressor.CreateHeader(fileHeader)
	if err != nil {
		return 0, err
	}

	numberOfBytes, err := compressed.Write(fileData)
	if err != nil {
		return 0, err
	}

	return numberOfBytes, nil
}

func ZipFile(filePath string, dest string) {
	calc := utils.GetRunTimeCalculator()
	calc.Start()

	fileData, fileInfo, err := readFileDataAndInfo(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fd, compressor, err := getCompressor(dest + "\\" + strings.Split(fileInfo.Name(), ".")[0] + ".zip")
	if err != nil {
		fmt.Println(err)
	}

	fileHeader, err := makeFileHeader(fileInfo)
	if err != nil {
		fmt.Println(err)
	}

	numberOfBytes, err := compress(fileHeader, fileData, compressor)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("File Compressed, Bytes Written %v\n", numberOfBytes)
	calc.End()

	defer fd.Close()
	defer compressor.Close()
}

type fileDetails struct {
	path     string
	fileInfo os.FileInfo
}

func readDir(path string, c chan fileDetails, running *int64) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		// is file a folder
		newPath := path + "\\" + file.Name()
		if file.IsDir() {
			atomic.AddInt64(running, 1)
			go readDir(newPath, c, running)
		} else {
			// if not write it to compress
			c <- fileDetails{path: newPath, fileInfo: file}
		}
	}

	atomic.AddInt64(running, -1)
}

func writeFile(fDetails fileDetails, fd *os.File, compressor *zip.Writer, originalPath string, parentFolder string) {
	// read file bytes
	fileData, err := ioutil.ReadFile(fDetails.path)
	if err != nil {
		fmt.Println("Error reading file", fDetails.path)
	}

	fileHeader, err := makeFileHeader(fDetails.fileInfo)
	if err != nil {
		fmt.Println(err)
	}

	relativePath := strings.Split(fDetails.path, originalPath)[1]
	fileHeader.Name = parentFolder + strings.Replace(relativePath, "\\", "/", -1)

	_, err = compress(fileHeader, fileData, compressor)
	if err != nil {
		fmt.Println(err)
	}
}

func ZipFolder(src string, dest string) {

	calc := utils.GetRunTimeCalculator()
	calc.Start()

	fmt.Printf("Processing... \n")

	c := make(chan fileDetails, 50)
	var running int64

	parts := strings.Split(src, "\\")
	foldername := parts[len(parts)-1]

	fd, compressor, err := getCompressor(dest + "\\" + foldername + ".zip")
	if err != nil {
		fmt.Println(err)
	}

	atomic.AddInt64(&running, 1)
	go readDir(src, c, &running)

	for file := range c {
		// write file
		writeFile(file, fd, compressor, src, foldername)
		if running == 0 {
			break
		}
	}

	fmt.Printf("Processing Complete \n")
	calc.End()

	defer fd.Close()
	defer compressor.Close()
}
