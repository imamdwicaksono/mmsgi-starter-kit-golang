package generators

var mainGoTemplate = `package main

import (
	"log"
	"net/http"

	"{{PROJECT}}/internal/routes"
)

func main() {
	router := routes.SetupRouter()

	// static files from ./public
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`

var routesTemplateBase = `package routes

import (
	"github.com/go-chi/chi/v5"
	"{{PROJECT}}/internal/controllers"
	"{{PROJECT}}/internal/middlewares"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.Logger)
	{{ROUTE_BODY}}
	return r
}`

var homeTemplate = `package controllers

import (
	"net/http"

	"github.com/mmsgi-library/mmsgi-starter-kit/pkg/templates"
)

func RenderTemplate(w http.ResponseWriter, name string) {
	templates.Render(w, name, nil)
}

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, "dashboard.html", nil)
}

func AboutIndex(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, "about.html", nil)
}
`

var loggerTemplate = `package middlewares

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
`

var dashboardTemplate = `{{define "dashboard"}}<h1>Dashboard</h1><p>Welcome to MMSGI Starter Kit</p>{{end}}`
var aboutTemplate = `{{define "about"}}<h1>About</h1><p>Starter Kit by MMSGI</p>{{end}}`

var defaultCSS = `body { font-family: Arial, sans-serif; padding: 20px; background: #f4f4f4; } .navbar{background:#333;color:#fff;padding:10px} .navbar a{color:#fff;margin-right:10px;text-decoration:none}`
var defaultJS = `console.log("MMSGI project loaded");`

// small bootstrap-ish css (very small subset)
var bootstrapCSS = `/* Lightweight bootstrap-like */ body{font-family:Arial,Helvetica,sans-serif;margin:0;padding:0;background:#f8f9fa} .container{max-width:960px;margin:24px auto;padding:16px;background:#fff;border-radius:6px}`

// tailwind fallback "compiled" minimal to keep look mostly same without npm build
var tailwindFallbackCSS = `/* tailwind fallback minimal */ body{font-family:system-ui,-apple-system,Segoe UI,Roboto,Helvetica,Arial,sans-serif;background:#f8fafc;margin:0} .container{max-width:1024px;margin:32px auto;padding:20px;background:#fff;border-radius:8px}`
var tailwindConfigJS = `module.exports = {
  content: ["./views/**/*.{html,tmpl}", "./src/**/*.{html,js}"],
  theme: { extend: {} },
  plugins: [],
};`
var tailwindPackageJSON = `{
  "name": "mmsgi-tailwind",
  "private": true,
  "devDependencies": {
    "tailwindcss": "^3.4.0",
    "postcss": "^8.4.0",
    "autoprefixer": "^10.4.0"
  },
  "scripts": {
    "build:css": "tailwindcss -i src/input.css -o public/css/style.css --minify"
  }
}`
var tailwindInputCSS = `@tailwind base;
@tailwind components;
@tailwind utilities;`

var readmeTemplate = `# {{PROJECT}}`
