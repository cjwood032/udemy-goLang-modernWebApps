package handlers

import (
	"github.com/cjwood032/go-course/pkg/config"
	"github.com/cjwood032/go-course/pkg/models"
	"github.com/cjwood032/go-course/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.ApplConfig
}

func NewRepo(a *config.ApplConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
