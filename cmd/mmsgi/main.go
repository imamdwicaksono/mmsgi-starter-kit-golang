package main

import (
	"fmt"
	"os"

	"github.com/mmsgi-library/mmsgi-starter-kit/internal/generators"
	"github.com/spf13/cobra"
)

func main() {
	var root = &cobra.Command{
		Use:   "mmsgi",
		Short: "MMSGI Starter Kit Generator",
	}

	var cmdNew = &cobra.Command{
		Use:   "new [project name]",
		Short: "Generate new MMSGI web project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			api, _ := cmd.Flags().GetBool("api")
			auth, _ := cmd.Flags().GetBool("auth")
			crud, _ := cmd.Flags().GetBool("crud")
			ui, _ := cmd.Flags().GetString("ui")

			opts := generators.GenerateOptions{
				UseAPI:  api,
				UseAuth: auth,
				UseCRUD: crud,
				UI:      ui,
			}

			if err := generators.GenerateProject(name, opts); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println("✔ Project created:", name)
		},
	}

	cmdNew.Flags().Bool("api", false, "include API mode (REST endpoints)")
	cmdNew.Flags().Bool("auth", false, "include authentication (session + login)")
	cmdNew.Flags().Bool("crud", false, "include CRUD scaffolding")
	cmdNew.Flags().String("ui", "tailwind", "choose UI (tailwind|bootstrap|none)")

	root.AddCommand(cmdNew)

	if err := root.Execute(); err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
}
