package server

import (
	"fmt"
	"log"
	"os"
)

type FileWriter struct {
	file *os.File
}

// Create New file writer
func NewFileWriter(file *os.File) *FileWriter {
	return &FileWriter{
		file: file,
	}
}

// Write implements the io.Writer interface
func (fw *FileWriter) Write(p []byte) (n int, err error) {
	return fw.file.Write(p)
}

// All Files are created or Opened under ~/.taskweaver directory
func CreateAndOpen(file string) (oFile *os.File, err error) {
	homeDir, _ := os.UserHomeDir()
	dirPath := homeDir + "/.taskweaver/"
	// Create the directory if it doesn't exist
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("Error Creating Directory -> %s", err)
	}
	// Open or create file
	File, err := os.OpenFile(dirPath+file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	return File, nil
}
