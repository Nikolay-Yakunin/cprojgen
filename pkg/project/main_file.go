package project

import (
	"path/filepath"
)

// createMainFile создаёт main.c в зависимости от типа проекта.
func createMainFile(root, projectType string) error {
	mainPath := filepath.Join(root, "src", "main.c")
	libHeaderPath := filepath.Join(root, "include", "mylib.h")
	libSrcPath := filepath.Join(root, "src", "mylib.c")

	mainBin := `#include <stdio.h>

int main(void) {
    printf("Hello, world!\n");
    return 0;
}`

	mainLib := `#include <stdio.h>
#include "mylib.h"

int main(void) {
    my_function();
    return 0;
}`

	libHeader := `#ifndef MYLIB_H
#define MYLIB_H

void my_function();

#endif`

	libSrc := `#include "mylib.h"
#include <stdio.h>

void my_function() {
    printf("Library function called!\n");
}`

	if projectType == "lib" {
		if err := createFile(libHeaderPath, libHeader); err != nil {
			return err
		}
		if err := createFile(libSrcPath, libSrc); err != nil {
			return err
		}
		return createFile(mainPath, mainLib)
	}

	return createFile(mainPath, mainBin)
}
