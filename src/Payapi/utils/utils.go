package utils

import (
	"net/http"
)

func ReceiveRequestAndParse(r *http.Request) (map[string]interface{}, error) {
	var err error
	var inputParams = make(map[string]interface{})
	for key, value := range r.Form {
		inputParams[key] = value[0]
	}
	return inputParams, err
}
