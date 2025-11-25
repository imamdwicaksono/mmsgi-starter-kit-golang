package generators

import (
	"os"
	"strings"
)

func CopyProjectFiles(name string) error {
	replacements := map[string]string{
		"{{PROJECT}}":         name,
		"{{MAIN_GO}}":         mainGoTemplate,
		"{{ROUTES}}":          routesTemplate,
		"{{CONTROLLER_HOME}}": homeTemplate,
		"{{LOGGER}}":          loggerTemplate,
		"{{DASHBOARD_HTML}}":  dashboardTemplate,
		"{{ABOUT_HTML}}":      aboutTemplate,
		"{{CSS}}":             defaultCSS,
		"{{JS}}":              defaultJS,
	}

	for path, content := range TemplateFiles {
		for k, v := range replacements {
			content = strings.ReplaceAll(content, k, v)
		}

		fullpath := name + "/" + path
		os.WriteFile(fullpath, []byte(content), 0644)
	}

	return nil
}
