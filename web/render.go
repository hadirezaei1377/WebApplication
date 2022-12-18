package main

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
