package checklistcontroller

import (
	// "Payapi/components/serviceerrors"
	// "Payapi/components/validators"
	// "Payapi/config/validationmaps/configvalidations"
	// "Payapi/templates/credit/applicationview"
	// "encoding/json"
	// "fmt"
	"net/http"
)

var Actions map[string]func(http.ResponseWriter, *http.Request) = map[string]func(http.ResponseWriter, *http.Request){}
