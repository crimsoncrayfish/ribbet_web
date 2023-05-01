package router

import (
	"context"
	"net/http"
	httpr "ribbet_web/common/http_r"
)

func (router router) HomeLoader(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request) error {

	router.l.Println("Loading home page")

	list, err := router.as.List()
	if err != nil {
		return err
	}
	return httpr.Template(w, http.StatusOK, "home-view", router.t, list)
}
