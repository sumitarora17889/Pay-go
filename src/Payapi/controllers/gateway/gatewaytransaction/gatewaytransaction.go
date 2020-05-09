package gatewaytransaction

import (
	// "Payapi/components/serviceerrors"
	// "Payapi/components/validators"
	// "Payapi/config/validationmaps/configvalidations"
	// "fmt"
	"net/http"
)

var Actions map[string]func(http.ResponseWriter, *http.Request) = map[string]func(http.ResponseWriter, *http.Request){}
