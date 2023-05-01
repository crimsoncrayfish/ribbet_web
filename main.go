package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	logr "ribbet_web/common/log_r"
	"ribbet_web/router"
	"ribbet_web/stores"
)

var (
	//go:embed templates
	templates embed.FS
	//go:embed assets
	assets embed.FS
)

func main() {
	l := logr.New("Ribbet Habbit Tracker")

	if err := run(l); err != nil {
		l.Printf("Failed with error %s", err)
		os.Exit(1)
	}
}

func run(l *log.Logger) error {
	// ====================================================================
	// Get config
	addr := ":4200"

	// ====================================================================
	// DB and Store setup
	l.Println("Setting up DB")
	activitystore, err := stores.ProvideActivityStore(l)
	if err != nil {
		return err
	}

	l.Println("DB setup complete")
	// ====================================================================
	// Template
	l.Println("Setting up templates")
	t, err := template.ParseFS(templates, "templates/*.go.html")
	if err != nil {
		return err
	}

	l.Println("Templates setup complete")
	// ====================================================================
	// Start Server
	l.Println("Setting up router and server")
	router := router.Init(l, activitystore, assets, t)
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	l.Printf("Listening on port %s\n", addr)
	return server.ListenAndServe()
}
