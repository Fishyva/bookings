package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Fishyva/bookings/internal/config"
	"github.com/Fishyva/bookings/internal/forms"
	"github.com/Fishyva/bookings/internal/models"
	"github.com/Fishyva/bookings/internal/render"
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

    render.RenderTemplate(w,r,"home.page.html",&models.TemplateData{})
}
// About is the about page handler
func (m *Repository)About(w http.ResponseWriter, r *http.Request){
    // perform  some logic
    stringMap := make(map[string]string)
    stringMap["test"] = "Hello, again."

    remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
    stringMap["remote_ip"] = remoteIP


    //send date the data to the template
    render.RenderTemplate(w,r,"about.page.html",&models.TemplateData{
        StringMap: stringMap,
    })

}
//Reservations Handler to display form 
func (m *Repository) Reservations(w http.ResponseWriter, r *http.Request) {
    
    var emptyReservation models.Reservation
    data := make(map[string]interface{})
    data["reservation"] = emptyReservation

   render.RenderTemplate(w,r,"make-reservation.page.html",&models.TemplateData{
    Form: forms.New(nil),
    Data: data,
   })
}
// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservations(w http.ResponseWriter, r *http.Request) {
   err := r.ParseForm()
   if err != nil {
    log.Println(err)
     return
   }
   reservation := models.Reservation{
    FirstName: r.Form.Get("first_name"),
    LastName: r.Form.Get("last_name"),
    Phone: r.Form.Get("phone"),
    Email: r.Form.Get("email"),
   }
   form := forms.New(r.PostForm)
    // form validation
  // form.Has("first_name", r)
  form.Required("first_name","last_name","email","phone")
  form.MinLength("first_name",3,r)
  form.MinLength("last_name",2,r)

   if !form.Valid() {
        data := make(map[string]interface{})
        data["reservation"] = reservation

        render.RenderTemplate(w,r,"make-reservation.page.html",&models.TemplateData{
            Form: form,
            Data: data,
           })
           return
   }
   
 }
//Generals is a handler for the html template Generals
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w,r,"generals.page.html",&models.TemplateData{})
 }
 //Majors is a handler for html template Majors
 func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w,r,"majors.page.html",&models.TemplateData{})
 }
// Availability is a method for rendering the search availibity template
 func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w,r,"availability.page.html",&models.TemplateData{})
 }
// PostAvailability submits our form request
 func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
    start := r.Form.Get("start")
    end := r.Form.Get("end")

    w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start,end)))
 }

 type jsonResponse struct {
     OK bool `json:"ok"`
     Message string  `json:"message"`
 }
// AvailabilityJSON handles request for availability and handles JSON response
 func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
   resp := jsonResponse{
       OK: true,
       Message: "Available!",
   }
   out,err := json.MarshalIndent(resp,"","     ")
   if err != nil {
       log.Println(err)
   }
 
   w.Header().Set("Content-Type","application/json")
   w.Write(out)

 }

 // Contact
 func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w,r,"contact.page.html",&models.TemplateData{})
 }

