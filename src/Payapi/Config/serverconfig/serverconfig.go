package serverconfig

import (
	"Payapi/controllers/credit/applicationcontroller"
	"net/http"
)

var AllowedHeaders []string = []string{"X-Requested-With"}
var AllowedOrigins []string = []string{"*"}
var AllowedMethods []string = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}

var Listeningport string = ":8080" //Listeningport should be a string starting with colon followed by port number

var AllowedIPs []string = []string{"*"}

var Routers map[string]map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]map[string]func(http.ResponseWriter, *http.Request){
	"credit": CreditRouter,
}

var CreditRouter map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"application": applicationcontroller.Actions,
}
