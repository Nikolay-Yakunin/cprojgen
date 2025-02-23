package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"cprojgen/pkg/project"
)

func main() {
	projectType := flag.String("type", "bin", "Тип проекта: bin или lib")
	projectName := flag.String("name", "project", "Название проекта")
	flag.Parse()

	if *projectType != "bin" && *projectType != "lib" {
		fmt.Println("Неверный тип проекта. Используйте 'bin' или 'lib'")
		os.Exit(1)
	}

	projectRoot := filepath.Join(".", *projectName)
	if err := os.MkdirAll(projectRoot, 0755); err != nil {
		fmt.Println("Ошибка создания директории проекта:", err)
		os.Exit(1)
	}

	if err := project.CreateProjectStructure(projectRoot, *projectType); err != nil {
		fmt.Println("Ошибка создания структуры проекта:", err)
		os.Exit(1)
	}

	// Инициализация Git
	cmd := exec.Command("git", "init", projectRoot)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Ошибка инициализации git-репозитория:", err)
	}

	fmt.Println("Проект успешно инициализирован!")
}
