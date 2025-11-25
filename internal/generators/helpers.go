package generators

import "strings"

// buildRoutesTemplate returns the routes.go content based on options
func buildRoutesTemplate(opts GenerateOptions, project string) string {
	routeBody := []string{
		`r.Get("/", controllers.HomeIndex)`,
		`r.Get("/about", controllers.AboutIndex)`,
	}
	if opts.UseAuth {
		routeBody = append(routeBody, `r.Get("/login", controllers.LoginForm)`, `r.Post("/login", controllers.LoginPost)`, `r.Get("/logout", controllers.Logout)`)
	}
	if opts.UseAPI || opts.UseCRUD {
		routeBody = append(routeBody, `
r.Route("/api", func(r chi.Router) {
    r.Get("/posts", controllers.PostsIndex)
    r.Get("/posts/{id}", controllers.PostShow)
    r.Post("/posts", controllers.PostCreate)
})`)
	}
	joined := strings.Join(routeBody, "\n\t")
	res := strings.ReplaceAll(routesTemplateBase, "{{ROUTE_BODY}}", joined)
	// replace project placeholder with actual module path
	res = strings.ReplaceAll(res, "{{PROJECT}}", project)
	return res
}
