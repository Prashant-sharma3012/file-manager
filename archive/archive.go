package archive

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

func ZipFolder(src string, dest string) {

}
