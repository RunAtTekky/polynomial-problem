package helper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func TakeInput(dir string) string {
	filePath := filepath.Join(dir, "input.json")
	fmt.Printf("Reading input.json from %q\n", filePath)

	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("Error reading input.json file")
	}

	return string(content)
}
