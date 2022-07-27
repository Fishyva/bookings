package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Fishyva/bookings/internal/config"
	"github.com/Fishyva/bookings/internal/models"
	"github.com/justinas/nosurf"
)
var functions = template.FuncMap{}
var app *config.AppConfig
//New Template Sets the config for the template package
func NewTemplate(a *config.AppConfig){
    app = a
}
func AddDefaultData(td *models.TemplateData,r *http.Request) *models.TemplateData {
    td.Flash = app.Session.PopString(r.Context(), "flash")
    td.Error = app.Session.PopString(r.Context(), "error")
    td.Warning = app.Session.PopString(r.Context(), "warning")
    td.CSRFToken = nosurf.Token(r)
    return td
}
// get the template cache from the app config
func RenderTemplate(w http.ResponseWriter,r *http.Request ,html string,td *models.TemplateData){
    var tc map[string]*template.Template
    if app.UseCache {
    // get the template cache from the app config
    tc = app.TemplateCache
    } else {
        tc,_ = CreateTemplateCache()
    }
      //getting our single template from template cache   
    myTemplate, ok := tc[html]
    if !ok {
        log.Fatal("could not get template from template cache")
    }
    buf := new(bytes.Buffer)
    td = AddDefaultData(td, r)
    _ = myTemplate.Execute(buf,td)
    _,err := buf.WriteTo(w) 
    if err != nil {
        fmt.Println("Error printing to browser",err)
    }
}
// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template,error) {
    
    myCache := map[string]*template.Template{}
    //step 1 need to get the files at a location
    pages,err := filepath.Glob("./templates/*.page.html")
    if err != nil {
        return myCache,err
    }
    // getting the indivisual pages from templates folder
    for _, page := range pages {
        name := filepath.Base(page)

        hs, err := template.New(name).Funcs(functions).ParseFiles(page)
        if err != nil {
            return myCache,err
        }
        // see if your layout matches 
        matches, err := filepath.Glob("./templates/*layout.html")
        if err != nil {
            return myCache,err
        }
        if len(matches) > 0 {
            hs, err = hs.ParseGlob("./templates/*layout.html")
            if err != nil {
                return myCache,err
            }
        }
        myCache[name] = hs

    }
    return myCache,nil
}