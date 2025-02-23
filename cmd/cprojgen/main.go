package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// createFile создаёт файл по указанному пути с заданным содержимым.
func createFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

// createDirectories создаёт все директории, указанные в срезе dirs относительно корня root.
func createDirectories(root string, dirs []string) error {
	for _, d := range dirs {
		fullPath := filepath.Join(root, d)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return err
		}
	}
	return nil
}

// createProjectStructure создает структуру проекта, исходники, Makefile и README.md.
func createProjectStructure(root, projectType string) error {
	// Определяем базовую структуру директорий
	dirs := []string{"src", "include", "build", "bin", "tests"}
	if err := createDirectories(root, dirs); err != nil {
		return err
	}

	// Создаем исходный файл main.c в каталоге src
	var mainContent string
	if projectType == "lib" {
		// Для библиотеки создаём заголовочный файл и исходник реализации функции
		headerContent := "#ifndef MYLIB_H\n#define MYLIB_H\n\nvoid my_function();\n\n#endif // MYLIB_H\n"
		if err := createFile(filepath.Join(root, "include", "mylib.h"), headerContent); err != nil {
			return err
		}
		// Файл с реализацией функции библиотеки
		libContent := "#include \"mylib.h\"\n\nvoid my_function() {\n    // TODO: Реализуйте функцию\n}\n"
		if err := createFile(filepath.Join(root, "src", "mylib.c"), libContent); err != nil {
			return err
		}
		// Файл main.c для тестового вызова библиотеки
		mainContent = "#include <stdio.h>\n#include \"mylib.h\"\n\nint main(void) {\n    printf(\"Library project compiled successfully.\\n\");\n    return 0;\n}\n"
	} else {
		// Для бинарного проекта
		mainContent = "#include <stdio.h>\n\nint main(void) {\n    printf(\"Hello, world!\\n\");\n    return 0;\n}\n"
	}
	if err := createFile(filepath.Join(root, "src", "main.c"), mainContent); err != nil {
		return err
	}

	// Создание Makefile, который поддерживает сборку под Linux и Windows
	makefileContent := `# Определяем компилятор
CC = gcc

# Определяем расширение для исполняемого файла: если Windows, добавляем .exe
ifeq ($(OS),Windows_NT)
	EXE_EXT = .exe
else
	EXE_EXT =
endif

# Директории
SRC_DIR    = src
INCLUDE_DIR= include
OBJ_DIR    = build
BIN_DIR    = bin

# Файлы
SOURCES  = $(wildcard $(SRC_DIR)/*.c)
OBJECTS  = $(patsubst $(SRC_DIR)/%.c, $(OBJ_DIR)/%.o, $(SOURCES))
TARGET   = $(BIN_DIR)/program$(EXE_EXT)

.PHONY: all setup clean

# Основная цель: сборка проекта
all: setup $(TARGET)
	@echo "Компиляция завершена успешно."

# Создание необходимых директорий
setup:
	@mkdir -p $(OBJ_DIR) $(BIN_DIR)

# Компиляция исходных файлов в объектные
$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c
	$(CC) -c $< -o $@ -I$(INCLUDE_DIR)

# Линковка объектных файлов в исполняемый файл
$(TARGET): $(OBJECTS)
	$(CC) $^ -o $@

# Очистка сгенерированных файлов
clean:
	@rm -rf $(OBJ_DIR) $(BIN_DIR)
	@echo "Очищено."
`
	if err := createFile(filepath.Join(root, "Makefile"), makefileContent); err != nil {
		return err
	}

	// Создание файла README.md с описанием проекта
	readmeContent := "# Шаблон проекта на C\n\n" +
		"Этот проект является шаблоном для разработки на языке C, который поддерживает компиляцию как под Linux, так и под Windows (при наличии соответствующей среды).\n\n" +
		"## Структура проекта\n\n" +
		"- **src/** — исходные файлы (.c).\n" +
		"- **include/** — заголовочные файлы (.h).\n" +
		"- **build/** — директория для объектных файлов (создаётся автоматически).\n" +
		"- **bin/** — директория для исполняемого файла (создаётся автоматически).\n" +
		"- **tests/** — тестовые файлы (если необходимо).\n" +
		"- **Makefile** — скрипт для сборки проекта.\n" +
		"- **README.md** — это описание проекта.\n\n" +
		"## Сборка\n\n" +
		"Для компиляции проекта выполните в терминале команду:\n\n" +
		"```sh\nmake\n```\n\n" +
		"После успешной сборки исполняемый файл появится в директории **bin**.\n\n" +
		"## Очистка\n\n" +
		"Чтобы удалить сгенерированные объектные файлы и исполняемый файл, выполните команду:\n\n" +
		"```sh\nmake clean\n```\n\n" +
		"## Запуск\n\n" +
		"Запустите программу, перейдя в директорию **bin** и запустив исполняемый файл (например, `./program` для Linux или `program.exe` для Windows).\n"
	if err := createFile(filepath.Join(root, "README.md"), readmeContent); err != nil {
		return err
	}

	return nil
}

func main() {
	// Флаги для указания типа проекта (bin или lib) и названия проекта
	projectType := flag.String("type", "bin", "Тип проекта: bin или lib")
	projectName := flag.String("name", "project", "Название проекта")
	flag.Parse()

	if *projectType != "bin" && *projectType != "lib" {
		fmt.Println("Неверный тип проекта. Используйте 'bin' или 'lib'")
		os.Exit(1)
	}

	// Создаем корневую директорию проекта
	if err := os.MkdirAll(*projectName, 0755); err != nil {
		fmt.Println("Ошибка создания директории проекта:", err)
		os.Exit(1)
	}

	// Создаем структуру проекта и все необходимые файлы
	if err := createProjectStructure(*projectName, *projectType); err != nil {
		fmt.Println("Ошибка создания структуры проекта:", err)
		os.Exit(1)
	}

	// Инициализируем Git-репозиторий
	cmd := exec.Command("git", "init", *projectName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Ошибка инициализации git-репозитория:", err)
	}

	fmt.Println("Проект успешно инициализирован!")
}
