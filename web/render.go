package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
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

/*
 function for rendering a template
 page is a template we want to render it
 partial means etc
*/
func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	// normally we dont define err type
	var t *template.Template
	// what template do i want to render?
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	// If the desired template already exists
	_, templateInMap := app.templateCache[templateToRender]

	// add or remove templates feature aoutomatically and withaout need to stop and restart layers
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

func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	// build partials
	var err error

	// build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
	}
	// partial for developer
	// add partial
	if len(partials) > 0 {
		// use in folder template the file
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", templateToRender)
	}
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

}
