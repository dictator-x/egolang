package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func TruncWriteFile(filename string, s string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer, s)
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func AppendStringToFile(filename string, s string) {
	var file *os.File
	var err error

	if checkFileIsExist(filename) {
		file, err = os.OpenFile(filename, os.O_APPEND, 0666)
		fmt.Println("File: ", filename, " exist")
	} else {
		file, err = os.Create(filename)
		fmt.Println("File: ", filename, " not exist")
	}

	defer file.Close()
	n, err := io.WriteString(file, s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("write %d bytes into file: %s.", n, filename)
}
