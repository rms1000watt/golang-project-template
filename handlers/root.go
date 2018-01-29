package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Path not found: ", r.URL.Path)
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
