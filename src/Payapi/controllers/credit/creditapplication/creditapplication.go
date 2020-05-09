package creditapplication

import (
	"Payapi/components/utils"
	"Payapi/components/validators"
	"Payapi/config/validationmaps/configvalidations"
	"Payapi/models/credit/creditapplicationmodel"
	"Payapi/templates/credit/creditapplicationtemplate"
	"net/http"
)

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
	output := make(map[string]interface{})
	if err.Httpcode == 200 && err.Responsecode == 200 {
		output, err = creditapplicationmodel.Createapplication(args)
	}
	out, err := creditapplicationtemplate.Createapplicationresponse(output)
	utils.OutResponse(w, err, out)
}

func details(w http.ResponseWriter, r *http.Request) {
}

func listing(w http.ResponseWriter, r *http.Request) {
}

func update(w http.ResponseWriter, r *http.Request) {
}
