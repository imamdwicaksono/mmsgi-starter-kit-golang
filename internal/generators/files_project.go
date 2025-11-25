package generators

var TemplateFiles = map[string]string{
	"go.mod":                         "module {{PROJECT}}\n\ngo 1.21\n",
	"cmd/server/main.go":             "{{MAIN_GO}}",
	"internal/routes/routes.go":      "{{ROUTES}}",
	"internal/controllers/home.go":   "{{CONTROLLER_HOME}}",
	"internal/middlewares/logger.go": "{{LOGGER}}",
	"views/dashboard.html":           "{{DASHBOARD_HTML}}",
	"views/about.html":               "{{ABOUT_HTML}}",
	"public/css/style.css":           "{{CSS}}",
	"public/js/app.js":               "{{JS}}",
}
