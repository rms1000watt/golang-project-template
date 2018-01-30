package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Person struct {
	ID   string
	Name string
	Age  int
}

var PersonMap = map[string]Person{
	"4571CA38-BF46-4FA7-A507-1D2E124B9F2E": Person{
		ID:   "4571CA38-BF46-4FA7-A507-1D2E124B9F2E",
		Name: "Ryan Smith",
		Age:  201,
	},
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting personHandler")
	defer log.Debug("Finished personHandler")

	if r.Method == http.MethodOptions {
		// Process headers and return
		return
	}

	switch r.Method {
	case http.MethodGet:
		PersonHandlerGET(w, r)
	case http.MethodPost:
		PersonHandlerPOST(w, r)
	default:
		log.Debug("Method not allowed: ", r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}
