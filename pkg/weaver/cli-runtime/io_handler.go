package cli_runtime

import "os"

func HandleError(err error) {

}

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
