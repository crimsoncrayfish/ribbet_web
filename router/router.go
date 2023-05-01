package router

import (
	"context"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	httpr "ribbet_web/common/http_r"
	activitystore "ribbet_web/stores/activity_store"
)

type router struct {
	l  *log.Logger
	as activitystore.ActivityStore
	t  *template.Template
}

func Init(l *log.Logger,
	as activitystore.ActivityStore,
	assets fs.FS,
	t *template.Template) http.Handler {

	app := httpr.NewApp()

	fs := http.FileServer(http.FS(assets))

	app.Handle(http.MethodGet, "/assets/*", func(
		ctx context.Context,
		w http.ResponseWriter,
		r *http.Request,
	) error {
		fs.ServeHTTP(w, r)
		return nil
	})

	router := router{
		l:  l,
		as: as,
		t:  t,
	}

	app.Handle(http.MethodGet, "/", router.HomeLoader)

	return app
}
