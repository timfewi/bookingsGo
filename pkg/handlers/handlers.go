package handlers

import (
	"net/http"

	"github.com/timfewi/bookingsGo/pkg/config"
	"github.com/timfewi/bookingsGo/pkg/models"
	"github.com/timfewi/bookingsGo/pkg/render"
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

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r

}

// Home is a function that returns the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

// About is a function that returns the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	StringMap["remote_ip"] = remoteIP
	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}
