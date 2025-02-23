package project

import (
	"os"
	"path/filepath"
)

// createDirectories создаёт каталоги.
func createDirectories(root string, dirs []string) error {
	for _, d := range dirs {
		fullPath := filepath.Join(root, d)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return err
		}
	}
	return nil
}

// createFile создаёт файл и записывает в него содержимое.
func createFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}
