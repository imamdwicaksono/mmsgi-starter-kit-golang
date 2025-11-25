package generators

var loginViewTemplate = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8"/>
  <title>Login - MMSGI</title>
  <link rel="stylesheet" href="/assets/css/style.css">
</head>
<body>
  <main>
    <h1>Login</h1>
    <form method="POST" action="/login">
      <label>Username <input name="username" /></label><br/>
      <label>Password <input name="password" type="password"/></label><br/>
      <button type="submit">Login</button>
    </form>
  </main>
</body>
</html>`

var authControllerTemplate = `package controllers

import (
	"net/http"
	"time"
)

const sessionCookieName = "mmsgi_session"

func LoginForm(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login.html")
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	// WARNING: demo only. Replace with DB-backed auth.
	if username == "admin" && password == "secret" {
		http.SetCookie(w, &http.Cookie{
			Name:    sessionCookieName,
			Value:   "valid-session-token",
			Expires: time.Now().Add(24 * time.Hour),
			Path:    "/",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   sessionCookieName,
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}`
