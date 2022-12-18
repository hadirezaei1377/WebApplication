package main

import "net/http"

// a function by pointer to application for access to app config
// call virtual terminal , its a handler
func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Hit the handler")
}
