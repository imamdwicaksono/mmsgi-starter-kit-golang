package templates

import (
    "embed"
    "html/template"
    "net/http"
)

//go:embed html/*.html
var TemplateFS embed.FS

var tpl = template.Must(template.ParseFS(TemplateFS, "html/*.html"))

func Render(w http.ResponseWriter, name string, data any) {
    tpl.ExecuteTemplate(w, name, data)
}
