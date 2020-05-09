package main

import (
	"Payapi/config/serverconfig"
	"fmt"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter().StrictSlash(true)
	/* Function to map service URLs to functionality */
	serverconfig.InitRoutes(route)

	/* Function to print all the URLs implemented by API*/
	erro := route.Walk(gorillaWalkFn)
	if erro != nil {
		fmt.Println(erro)
	}

	err := serverconfig.LoadserverConfig(route) /*Set listening port, Allow heads, Allow origins, Allow Methods. Config stored in Config/serverconfig/serverconfig.go*/
	if err != "" {
		fmt.Println(err)
		os.Exit(1)
	}
}

func gorillaWalkFn(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	path, err := route.GetPathTemplate()
	fmt.Println(path)
	return err
}
