package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

const version = "1.0.0"
const cssVersion = "1"

// configuration info for app
type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

// reciever for the various parts of app
type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

// create web server
// its have a reciever named app that point to application
func (app *application) serve() error {
	// srv get your values from server in http by &
	srv := &http.Server{
		// address
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
		// The amount of time that the connection will be closed automatically if the user does not use the connection.
		IdleTimeout: 30 * time.Second,
		//  covers the time from when the connection is accepted to when the request body is fully read
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		// the maximum duration before timing out writes of the respons
		WriteTimeout: 5 * time.Second,
	}

	app.infoLog.Println("Starting HTTP server in %s mode on port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()
}

func main() {
	var cfg config

	// enter cmd into config variable by port

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application enviornment {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to api")

	flag.Parse()

	// getting secret key for credit card info by clients

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	// log in

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// shortfile is used for possible problems

	// map for template cashe
	tc := make(map[string]*template.Template)

	// app variable for asign values and we use a reference to application by &
	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
	}

	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
