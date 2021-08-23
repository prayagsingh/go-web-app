package handlers

import (
	"net/http"

	"github.com/prayagsingh/go-web-app/pkg/config"
	"github.com/prayagsingh/go-web-app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handler
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.html")
}

// About is the handler for the about page
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.page.html")
}
