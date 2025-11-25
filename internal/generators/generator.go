package generators

import (
	"fmt"
	"os"
)

func GenerateProject(name string) error {
	fmt.Println("Generating project:", name)

	// create folders
	folders := []string{
		name,
		name + "/cmd/server",
		name + "/internal/controllers",
		name + "/internal/routes",
		name + "/internal/middlewares",
		name + "/views",
		name + "/public/css",
		name + "/public/js",
	}

	for _, f := range folders {
		if err := os.MkdirAll(f, 0755); err != nil {
			return err
		}
	}

	// copy template files
	if err := CopyProjectFiles(name); err != nil {
		return err
	}

	return nil
}
