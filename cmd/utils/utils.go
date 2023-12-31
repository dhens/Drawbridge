package utils

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
)

// Save file if it does not already exist.
func SaveFile(fileName string, fileContents string, relativePath string) error {
	fullFilePath := fmt.Sprintf("%s/%s", relativePath, fileName)
	// Verify the file doesn't exist before opening.
	_, err := os.Open(fullFilePath)
	if !errors.Is(err, os.ErrNotExist) {
		log.Printf("Skipping existing save file: %s\n", fullFilePath)
		return nil
	}

	f, err := os.Create(fullFilePath)
	if err != nil {
		return fmt.Errorf("error creating file: %s/%s", fileName, relativePath)
	}
	defer f.Close()

	_, err = f.WriteString(fileContents)
	if err != nil {
		return fmt.Errorf("error writing file contents: %s", err)
	}

	return nil
}

func FileExists(pathWithFilename string) bool {
	_, err := os.Open(pathWithFilename)
	return !errors.Is(err, os.ErrNotExist)
}

func ReadFile(pathWithFilename string) *[]byte {
	file, err := os.ReadFile(pathWithFilename)
	if !errors.Is(err, os.ErrNotExist) {
		return &file
	}
	return nil
}

func Sloggerf(level string, message string, args ...any) {
	switch level {
	case "info":
		slog.Info(fmt.Sprintf(message, args...))
	case "debug":
		slog.Debug(fmt.Sprintf(message, args...))
	}

}
