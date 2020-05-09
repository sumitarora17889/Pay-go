package utils

import (
	"Payapi/components/serviceerrors"
	"encoding/json"
	"fmt"
	"net/http"
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

func OutResponse(w http.ResponseWriter, err serviceerrors.Error, output interface{}) {
	bolB, _ := json.Marshal(output)
	fmt.Fprintf(w, string(bolB))
}
