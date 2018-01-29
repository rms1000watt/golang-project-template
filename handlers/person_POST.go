package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func PersonHandlerPOST(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting PersonHandlerPOST")
	defer log.Debug("Finished PersonHandlerPOST")

	// personIn, err := data.GetPersonIn(r)
	// if err != nil {
	// 	log.Error(err)
	// 	http.Error(w, helpers.ErrorJSON(err.Error()), http.StatusInternalServerError)
	// 	return
	// }
	// log.Debug(personIn)

	// personOut := data.PersonOut{}

	// // Developer do stuff here

	// helpers.WriteJSON(personOut, w)

}
