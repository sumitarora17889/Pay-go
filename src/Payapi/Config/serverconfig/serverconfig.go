package serverconfig

import (
	"Payapi/config/routes"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var AllowedHeaders []string = []string{"X-Requested-With"}
var AllowedOrigins []string = []string{"*"}
var AllowedMethods []string = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}

var Listeningport string = ":8080" //Listeningport should be a string starting with colon followed by port number

var AllowedIPs []string = []string{"*"}

func LoadserverConfig(router *mux.Router) string {
	//load server config
	//Set listening port
	//Allowed heads
	//Allowed origin
	//Allowed Methods
	//security config
	AllowedHeaders := handlers.AllowedHeaders(AllowedHeaders)
	AllowedOrigins := handlers.AllowedHeaders(AllowedOrigins)
	AllowedMethods := handlers.AllowedHeaders(AllowedMethods)

	err := http.ListenAndServe(Listeningport, handlers.CORS(AllowedHeaders, AllowedOrigins, AllowedMethods)(LowerCaseURI(router)))
	if err != nil {
		return "Count not start the server because of following err " + err.Error()
	}
	return ""
}

func InitRoutes(route *mux.Router) {
	for routername, routers := range routes.Routers {
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
