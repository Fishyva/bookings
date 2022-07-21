package handlers

import (
	"net/http"
    "github.com/Fishyva/bookings/pkg/render"
	"github.com/Fishyva/bookings/pkg/config"
	"github.com/Fishyva/bookings/pkg/models"
)


var Repo *Repository
// Repository is the repository type 
type Repository  struct {
    App *config.AppConfig
}
// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
    return &Repository {
        App: a,
    }
}
// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository){
    Repo = r
}
// Home is the home page handler
func (m *Repository)Home(w http.ResponseWriter, r *http.Request) {
    remoteIP := r.RemoteAddr

    m.App.Session.Put(r.Context(),"remote_ip",remoteIP)

    render.RenderTemplate(w,"home.page.html",&models.TemplateData{})
}
// About is the about page handler
func (m *Repository)About(w http.ResponseWriter, r *http.Request){
    // perform  some logic
    stringMap := make(map[string]string)
    stringMap["test"] = "Hello, again."

    remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
    stringMap["remote_ip"] = remoteIP


    //send date the data to the template
    render.RenderTemplate(w,"about.page.html",&models.TemplateData{
        StringMap: stringMap,
    })

}
//Reservations Handler to display form 
func (m *Repository) Reservations(w http.ResponseWriter, r *http.Request) {
   render.RenderTemplate(w,"make-reservation.page.html",&models.TemplateData{})
}
//Generals is a handler for the html template Generals
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w,"generals.page.html",&models.TemplateData{})
 }
 //Majors is a handler for html template Majors
 func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w,"majors.page.html",&models.TemplateData{})
 }
// Availability is a method for rendering the search availibity template
 func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w,"availability.page.html",&models.TemplateData{})
 }

 // Contact
 func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w,"contact.page.html",&models.TemplateData{})
 }