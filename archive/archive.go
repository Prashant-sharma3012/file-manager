package archive

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ZipFile(filePath string, dest string) {

	// read file bytes
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
	}

	compressedFilePath := dest + "\\" + strings.Split(fileInfo.Name(), ".")[0] + ".zip"
	fd, err := os.Create(compressedFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()

	fileHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		fmt.Println(err)
	}

	// fileHeader.Name = fileInfo.Name()
	fileHeader.Method = zip.Deflate

	compressor := zip.NewWriter(fd)
	compressed, err := compressor.CreateHeader(fileHeader)
	if err != nil {
		fmt.Println(err)
	}
	defer compressor.Close()

	numberOfBytes, err := compressed.Write(fileData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("File Compressed, Bytes Written %v\n", numberOfBytes)
}

func ZipFolder(src string, dest string) {

}
