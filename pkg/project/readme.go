package project

import (
	"path/filepath"
)

// createReadme создаёт README.md.
func createReadme(root, projectType string) error {
	readmePath := filepath.Join(root, "README.md")

	readmeContent := "# C Project Template\n\n" +
		"## Structure\n" +
		"- **src/** - Source files\n" +
		"- **include/** - Header files\n" +
		"- **build/** - Object files\n" +
		"- **bin/** - Executables\n" +
		"- **tests/** - Test files\n" +
		"- **Makefile** - Build script\n" +
		"- **README.md** - This file\n\n" +
		"## Build\n" +
		"`make`\n\n" +
		"## Clean\n" +
		"`make clean`\n"

	if projectType == "lib" {
		readmeContent += "## Library Build\n" +
			"- Static library: `make static`\n" +
			"- Shared library: `make shared`\n"
	}

	return createFile(readmePath, readmeContent)
}
