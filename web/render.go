package main

import (
	"embed"
	"fmt"
	"net/http"
	"text/template"
)

// turns website code into the interactive pages

// passing some info to our template. all of things we need in back end of app
type templateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	// for each of data not defeined above
	Data map[string]interface{}
	// a secure random token
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

// passing functions to templaet

var functions = template.FuncMap{}

// template file system
var templateFS embed.FS

// create function to put pass data to template data in output by pointer
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	return td
}

// td is any info in struct template data that we will add

// function for rendering a template
// page is a template we want to render it
// partial means etc
func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.tmpl", page)

	_, templateInMap := app.templateCache[templateToRender]

	if app.config.env == "production" && templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}
