package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateProject(name string, opts GenerateOptions) error {
	fmt.Println("→ Creating project:", name)

	paths := []string{
		name,
		filepath.Join(name, "cmd", "server"),
		filepath.Join(name, "internal", "controllers"),
		filepath.Join(name, "internal", "routes"),
		filepath.Join(name, "internal", "middlewares"),
		filepath.Join(name, "views"),
		filepath.Join(name, "public", "css"),
		filepath.Join(name, "public", "js"),
	}

	// extra if API/CRUD/Auth
	if opts.UseAPI || opts.UseCRUD {
		paths = append(paths, filepath.Join(name, "internal", "models"))
	}
	if opts.UseAuth {
		paths = append(paths, filepath.Join(name, "internal", "auth"))
	}

	for _, p := range paths {
		if err := os.MkdirAll(p, 0755); err != nil {
			return err
		}
	}

	fmt.Println("→ Writing files...")
	if err := CopyProjectFiles(name, opts); err != nil {
		return err
	}

	// create example CRUD if requested
	if opts.UseCRUD {
		if err := generateExampleCRUD(name); err != nil {
			return err
		}
	}

	fmt.Println("✔ Done!")
	return nil
}
