package handlers

import (
	"encoding/json"
	"log"
	"main/validators"
	"net/http"
)

func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		HandleError(w, err.Error(), http.StatusBadRequest)
	}

	// log recieved
	log.Printf("Received: %+v\n", person)

	err = validators.ValidateCredentials(person.Email, person.Password)
	if err != nil {
		HandleError(w, err.Error(), http.StatusBadRequest)
	}

	ForwardInput(person)

	if err == nil {
		HandleSuccess(w, "Successful signin!", http.StatusAccepted)
	}
}