package eligibilitycontroller

import (
	"Payapi/components/validators"
	"Payapi/config/validationmaps/configvalidations"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var creditrouters map[string]func(http.ResponseWriter, *http.Request) = map[string]func(http.ResponseWriter, *http.Request){
	"get": get,
	"set": set,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.CurrentRoute(r)
	for action, handler := range creditrouters {
		router.PathPrefix(action + "/").HandlerFunc(handler)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	validators.BasicValidation(r, configvalidations.CreateApplicationMap)

}

func set(w http.ResponseWriter, r *http.Request) {
	validators.BasicValidation(r, configvalidations.CreateApplicationMap)
}

func ollo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
