package generators

import (
	"os"
	"path/filepath"
)

func generateExampleCRUD(root string) error {
	// model
	model := `package models

type Post struct {
	ID    int64  ` + "`json:\"id\"`" + `
	Title string ` + "`json:\"title\"`" + `
	Body  string ` + "`json:\"body\"`" + `
}
`
	if err := os.WriteFile(filepath.Join(root, "internal", "models", "post.go"), []byte(model), 0644); err != nil {
		return err
	}

	// controller (simple in-memory)
	ctrl := `package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	ID    int64  ` + "`json:\"id\"`" + `
	Title string ` + "`json:\"title\"`" + `
	Body  string ` + "`json:\"body\"`" + `
}

var (
	postStore = make(map[int64]*Post)
	postIDSeq int64 = 1
	postMu    sync.Mutex
)

func PostsIndex(w http.ResponseWriter, r *http.Request) {
	postMu.Lock()
	defer postMu.Unlock()
	list := make([]*Post, 0, len(postStore))
	for _, p := range postStore {
		list = append(list, p)
	}
	json.NewEncoder(w).Encode(list)
}

func PostShow(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	postMu.Lock()
	defer postMu.Unlock()
	if p, ok := postStore[id]; ok {
		json.NewEncoder(w).Encode(p)
		return
	}
	http.NotFound(w, r)
}

func PostCreate(w http.ResponseWriter, r *http.Request) {
	var p Post
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	postMu.Lock()
	defer postMu.Unlock()
	p.ID = postIDSeq
	postIDSeq++
	postStore[p.ID] = &p
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(p)
}
`
	if err := os.WriteFile(filepath.Join(root, "internal", "controllers", "crud_post.go"), []byte(ctrl), 0644); err != nil {
		return err
	}
	return nil
}
