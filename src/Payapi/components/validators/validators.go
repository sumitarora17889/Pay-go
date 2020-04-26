package validators

import (
	"Payapi/components/serviceerrors"
	"Payapi/components/validators/datetimevalidator"
	"Payapi/components/validators/floatvalidator"
	"Payapi/components/validators/numbervalidator"
	"Payapi/components/validators/stringvalidator"
	"net/http"
	"strings"
)

func BasicValidation(r *http.Request, rules map[string]map[string]string) (map[string]interface{}, serviceerrors.Error) {
	keymap := make(map[string]interface{})
	var val interface{}
	var key string
	var keyrule map[string]string
	var err serviceerrors.Error
	r.ParseForm()
	for key, keyrule = range rules {
		formval := r.Form[key]
		if len(formval) <= 0 {
			if keyrule["Required"] == "True" {
				return nil, serviceerrors.ParamMissing(key, "PARAMMISSING")
			}
			keymap[key] = nil
		} else {
			value := strings.Join(formval, ",")
			if keyrule["Type"] == "Number" {
				val, err = numbervalidator.Validation(key, value, keyrule)
			} else if keyrule["Type"] == "Float" {
				val, err = floatvalidator.Validation(key, value, keyrule)
			} else if keyrule["Type"] == "DateTime" {
				val, err = datetimevalidator.Validation(key, value, keyrule)
			} else {
				val, err = stringvalidator.Validation(key, value, keyrule)
			}
			if err.Httpcode != 200 || err.Responsecode != 200 {
				return nil, err
			} else {
				keymap[key] = val
			}
		}
	}
	return keymap, err
}
