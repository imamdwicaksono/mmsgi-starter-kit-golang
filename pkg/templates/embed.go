package templates

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed html/*.html
var TemplateFS embed.FS

var templates *template.Template

func init() {
	tplFS, err := fs.Sub(TemplateFS, "html")
	if err != nil {
		// if embed not found, templates will be nil, generated projects will have their own views
		return
	}
	templates = template.Must(template.ParseFS(tplFS, "*.html"))
}

// Render simple wrapper. If embedded templates not present (nil), try to serve from file-based views.
func Render(w http.ResponseWriter, name string, data any) {
	if templates != nil {
		_ = templates.ExecuteTemplate(w, name, data)
		return
	}
	// fallback: try to read from local views folder
	t, err := template.ParseFiles("views/" + name)
	if err != nil {
		http.Error(w, "template error: "+err.Error(), 500)
		return
	}
	_ = t.Execute(w, data)
}
