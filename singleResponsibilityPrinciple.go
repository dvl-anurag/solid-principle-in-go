package main

import (
	"fmt"
	"io/ioutil"
)

// FileReader is responsible for reading data from a file
type FileReader struct {
	filePath string
}

// NewFileReader creates a new FileReader instance
func NewFileReader(filePath string) *FileReader {
	return &FileReader{filePath: filePath}
}

// ReadContent reads the content of the file
func (fr *FileReader) ReadContent() (string, error) {
	content, err := ioutil.ReadFile(fr.filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ConsolePrinter is responsible for printing content to the console
type ConsolePrinter struct{}

// PrintContent prints the content to the console
func (cp *ConsolePrinter) PrintContent(content string) {
	fmt.Println("Content:", content)
}
