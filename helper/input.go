package helper

import (
	"log"
	"os"
	"path/filepath"
)

func TakeInput(dir string) string {
	filePath := filepath.Join(dir, "input.json")

	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("Error reading input.json file")
	}

	return string(content)
}
