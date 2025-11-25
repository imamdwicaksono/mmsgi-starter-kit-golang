package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CopyProjectFiles(root string, opts GenerateOptions) error {
	replacements := defaultReplacements(root)

	// Add conditional templates
	if opts.UseAuth {
		replacements["{{LOGIN_HTML}}"] = loginViewTemplate
		replacements["{{AUTH_CONTROLLER}}"] = authControllerTemplate
		TemplateFiles["views/login.html"] = "{{LOGIN_HTML}}"
		TemplateFiles["internal/controllers/auth.go"] = "{{AUTH_CONTROLLER}}"
	}
	if opts.UseAPI || opts.UseCRUD {
		// Add routes for API (routesTemplate includes conditional markers)
		// We'll inject modified routes template below
	}
	// UI choices: if tailwind or bootstrap, add assets
	if opts.UI == "bootstrap" {
		replacements["{{CSS}}"] = bootstrapCSS // override default CSS
	} else if opts.UI == "tailwind" {
		// include a small compiled tailwind fallback (so it works without npm)
		replacements["{{CSS}}"] = tailwindFallbackCSS
		// also create package.json and tailwind config as helper files in root for user's use
		TemplateFiles["tailwind.config.js"] = tailwindConfigJS
		TemplateFiles["package.json"] = tailwindPackageJSON
		TemplateFiles["src/input.css"] = tailwindInputCSS
	} else {
		// leave default CSS
	}

	// Build routes content with options
	routes := buildRoutesTemplate(opts, root)
	replacements["{{ROUTES}}"] = routes

	// Fill in other replacements
	replacements["{{MAIN_GO}}"] = mainGoTemplate
	replacements["{{CONTROLLER_HOME}}"] = homeTemplate
	replacements["{{LOGGER}}"] = loggerTemplate
	replacements["{{DASHBOARD_HTML}}"] = dashboardTemplate
	replacements["{{ABOUT_HTML}}"] = aboutTemplate
	replacements["{{JS}}"] = defaultJS
	// default CSS if not overridden
	if _, ok := replacements["{{CSS}}"]; !ok {
		replacements["{{CSS}}"] = defaultCSS
	}

	for relPath, tpl := range TemplateFiles {
		content := tpl
		for k, v := range replacements {
			content = strings.ReplaceAll(content, k, v)
		}
		target := filepath.Join(root, relPath)
		dir := filepath.Dir(target)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		if err := os.WriteFile(target, []byte(content), 0644); err != nil {
			return fmt.Errorf("write %s: %w", target, err)
		}
	}

	// If UI == tailwind, also write package.json and config and src files
	if opts.UI == "tailwind" {
		if err := os.MkdirAll(filepath.Join(root, "src"), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(root, "tailwind.config.js"), []byte(tailwindConfigJS), 0644); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(root, "package.json"), []byte(tailwindPackageJSON), 0644); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(root, "src", "input.css"), []byte(tailwindInputCSS), 0644); err != nil {
			return err
		}
	}

	// Simple README
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte(fmt.Sprintf(readmeTemplate, root)), 0644); err != nil {
		return err
	}

	return nil
}

func defaultReplacements(root string) map[string]string {
	return map[string]string{
		"{{PROJECT}}": root,
	}
}
