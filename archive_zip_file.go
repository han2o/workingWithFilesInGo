/*
* zip压缩
 */
package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
)

func main() {
	outFile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)

	var fileToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}
	fmt.Println(fileToArchive)

	for _, file := range fileToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
