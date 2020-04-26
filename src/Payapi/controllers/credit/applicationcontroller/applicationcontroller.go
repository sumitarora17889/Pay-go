package applicationcontroller

import (
	"Payapi/components/serviceerrors"
	"Payapi/components/validators"
	"Payapi/config/validationmaps/configvalidations"
	"Payapi/templates/credit/applicationview"
	"encoding/json"
	"fmt"
	"net/http"
)

func OutResponse(w http.ResponseWriter, err serviceerrors.Error, output map[string]interface{}) {
	bolB, _ := json.Marshal(err)
	fmt.Fprintf(w, string(bolB))
}

var Actions map[string]func(http.ResponseWriter, *http.Request) = map[string]func(http.ResponseWriter, *http.Request){
	"create":  create,
	"details": details,
	"listing": listing,
	"update":  update,
}

// func Controller(w http.ResponseWriter, r *http.Request) {
// 	route := mux.CurrentRoute(r)
// 	pathTemplate, err := route.GetPathTemplate()
// 	if err == nil {
// 		fmt.Fprintf(w, pathTemplate)
// 	}
// 	// for action, handler := range creditrouters {
// 	// router.PathPrefix("create/").HandlerFunc(create)
// 	// }
// }

func create(w http.ResponseWriter, r *http.Request) {
	args, err := validators.BasicValidation(r, configvalidations.CreateApplicationMap)
	// OutResponse(w, err, args)
	// for key, value := range args {
	// 	fmt.Fprintf(w, key+":\t%v\n", value)
	// }
	out1, _ := json.Marshal(applicationview.Createemployee())
	output, _ := json.Marshal(args)
	fmt.Fprintf(w, string(output))
	fmt.Fprintf(w, "Hurray")
	fmt.Fprintf(w, err.Errortype)
	fmt.Fprintf(w, "\n"+string(out1))
}

func details(w http.ResponseWriter, r *http.Request) {
}

func listing(w http.ResponseWriter, r *http.Request) {
}

func update(w http.ResponseWriter, r *http.Request) {
}
