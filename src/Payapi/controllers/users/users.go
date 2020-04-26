package users

import (
	"Payapi/controllers/credit/applicationcontroller"
	"net/http"

	"github.com/gorilla/mux"
)

var creditrouters map[string]func(http.ResponseWriter, *http.Request) = map[string]func(http.ResponseWriter, *http.Request){
	"users": applicationcontroller.Handler,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.CurrentRoute(r)
	for controller, handler := range creditrouters {
		router.PathPrefix(controller + "/").HandlerFunc(handler)
	}
}
