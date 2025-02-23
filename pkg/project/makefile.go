package project

import (
	"path/filepath"
)

// createMakefile создаёт Makefile в зависимости от типа проекта.
func createMakefile(root, projectType string) error {
	makefilePath := filepath.Join(root, "Makefile")

	makefileBin := `CC = gcc
SRC_DIR = src
OBJ_DIR = build
BIN_DIR = bin
TARGET = $(BIN_DIR)/program
SOURCES = $(wildcard $(SRC_DIR)/*.c)
OBJECTS = $(patsubst $(SRC_DIR)/%.c, $(OBJ_DIR)/%.o, $(SOURCES))

all: setup $(TARGET)

setup:
	mkdir -p $(OBJ_DIR) $(BIN_DIR)

$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c
	$(CC) -c $< -o $@

$(TARGET): $(OBJECTS)
	$(CC) $^ -o $@

clean:
	rm -rf $(OBJ_DIR) $(BIN_DIR)
`

	makefileLib := `CC = gcc
SRC_DIR = src
OBJ_DIR = build
BIN_DIR = bin
LIB_DIR = lib
INCLUDE_DIR = include

STATIC_LIB = $(LIB_DIR)/mylib.a
SHARED_LIB = $(LIB_DIR)/mylib.so
SOURCES = $(wildcard $(SRC_DIR)/*.c)
OBJECTS = $(patsubst $(SRC_DIR)/%.c, $(OBJ_DIR)/%.o, $(SOURCES))

all: setup static shared

setup:
	mkdir -p $(OBJ_DIR) $(BIN_DIR) $(LIB_DIR)

$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c
	$(CC) -c $< -o $@ -I$(INCLUDE_DIR)

static: $(OBJECTS)
	ar rcs $(STATIC_LIB) $(OBJECTS)

shared: $(OBJECTS)
	$(CC) -shared -o $(SHARED_LIB) $(OBJECTS)

clean:
	rm -rf $(OBJ_DIR) $(LIB_DIR)
`

	content := makefileBin
	if projectType == "lib" {
		content = makefileLib
	}

	return createFile(makefilePath, content)
}
