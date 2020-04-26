package utils

import (
	"Payapi/config/serverconfig"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func ReceiveRequestAndParse(r *http.Request) (map[string]interface{}, error) {
	//Test
	var err error
	var inputParams = make(map[string]interface{})
	for key, value := range r.Form {
		inputParams[key] = value[0]
	}
	return inputParams, err
}
func LoadserverConfig(router *mux.Router) string {
	AllowedHeaders := handlers.AllowedHeaders(serverconfig.AllowedHeaders)
	AllowedOrigins := handlers.AllowedHeaders(serverconfig.AllowedOrigins)
	AllowedMethods := handlers.AllowedHeaders(serverconfig.AllowedMethods)

	err := http.ListenAndServe(serverconfig.Listeningport, handlers.CORS(AllowedHeaders, AllowedOrigins, AllowedMethods)(LowerCaseURI(router)))
	if err != nil {
		return "Count not start the server because of following err " + err.Error()
	}
	return ""
}

func InitRoutes(route *mux.Router) {
	for routername, routers := range serverconfig.Routers {
		for controllername, actions := range routers {
			for actionname, action := range actions {
				route.PathPrefix("/" + routername + "/" + controllername + "/" + actionname + "/").HandlerFunc(Middleware(action))
			}
		}
	}
	route.PathPrefix("/").HandlerFunc(notFoundHandler)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is error handler for path"+r.URL.Path[1:])
}
func LowerCaseURI(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func Middleware(action func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Executing Middleware start\n") //Add anything required for all services including logging requests and responses
		action(w, r)
		fmt.Fprintf(w, "Concluding Middleware")
	}

}
