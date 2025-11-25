package generators

import (
	"fmt"
	"os"
)

func GenerateProject(name string) error {
	fmt.Println("→ Creating project:", name)

	paths := []string{
		name,
		name + "/cmd/server",
		name + "/internal/controllers",
		name + "/internal/routes",
		name + "/internal/middlewares",
		name + "/views",
		name + "/public/css",
		name + "/public/js",
	}

	for _, p := range paths {
		os.MkdirAll(p, 0755)
	}

	fmt.Println("→ Writing files...")
	if err := CopyProjectFiles(name); err != nil {
		return err
	}

	fmt.Println("✔ Done!")
	return nil
}
