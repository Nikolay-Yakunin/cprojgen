package project

import (

)

// CreateProjectStructure создаёт директории, файлы и Makefile.
func CreateProjectStructure(root, projectType string) error {
	dirs := []string{"src", "include", "build", "bin", "tests"}
	if err := createDirectories(root, dirs); err != nil {
		return err
	}

	if err := createMainFile(root, projectType); err != nil {
		return err
	}

	if err := createMakefile(root, projectType); err != nil {
		return err
	}

	if err := createReadme(root, projectType); err != nil {
		return err
	}

	return nil
}
