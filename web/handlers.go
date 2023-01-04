package main

import "net/http"

/*
 a function by pointer to application for access to app config
 call virtual terminal , its a handler
*/

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	/*
	 app.infoLog.Println("Hit the handler")
	 instead of showing in terminal by above code, render the template
	*/
	if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
	// refer ro routes file and put it there
}
