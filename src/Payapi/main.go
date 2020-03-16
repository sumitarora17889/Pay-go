package main

import (
	"fmt"
	"net/http"
	"os"
	"utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//load server config
	//Set listening port
	//Allowed heads
	//Allowed origin
	//Allowed Methods
	//security config
	var err error
	var environment = os.Args[1]
	err = utils.LoadDefaultConfig(environment)
	if err != nil {
		fmt.Println("Could not load the configuration")
		os.Exit(1)
	}
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	router := mux.NewRouter().StrictSlash(true)

	port := utils.GetString("serviceListeningPort")
	if port == "" {
		fmt.Println("Port is not defined in configuration ...!!!")
		os.Exit(1)
	}
	err = http.ListenAndServe(port, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
	if err != nil {
		fmt.Println("Count not start the server because of following err " + err.Error())
		return
	}
}
