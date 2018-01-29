package handlers

import (
	"net/http"

	"github.com/rms1000watt/belly"
	log "github.com/sirupsen/logrus"
)

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting personHandler")
	defer log.Debug("Finished personHandler")

	if r.Method == http.MethodOptions {
		// Process headers and return
		return
	}

	switch r.Method {
	case http.MethodGet:
		belly.HandleMiddlewares(PersonHandlerGET, belly.MiddlewareNoCache)(w, r)
	case http.MethodPost:
		belly.HandleMiddlewares(PersonHandlerPOST, belly.MiddlewareNoCache)(w, r)
	default:
		log.Debug("Method not allowed: ", r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}
