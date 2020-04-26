package main

import (
	"Payapi/components/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main() {
	route := mux.NewRouter().StrictSlash(true)
	//load server config
	//Set listening port
	//Allowed heads
	//Allowed origin
	//Allowed Methods
	//security config
	utils.InitRoutes(route)

	erro := route.Walk(gorillaWalkFn)
	if erro != nil {
		fmt.Println(erro)
	}
	/*Set listening port, Allow heads, Allow origins, Allow Methods. Config stored in Config/serverconfig/serverconfig.go*/
	err := utils.LoadserverConfig(route)
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
