package validation

import (
	"datetimevalidator"
	"floatvalidator"
	"numbervalidator"
	"serviceerror"
	"stringvalidator"
)

func BasicValidation(inputParams map[string]interface{}, rules map[string]map[string]string) (map[string]interface{}, serviceerror.Error) {
	var keymap map[string]interface{}
	var key string
	var keyrule map[string]string
	var value interface{}
	var err serviceerror.Error
	for key, keyrule = range rules {
		value = inputParams[key]
		if value == "" {
			if keyrule["Required"] == "True" {
				return nil, serviceerror.ParamMissing("PARAMMISSING")
			} else {
				keymap[key] = nil
				continue
			}
		}
		if keyrule["Type"] == "Number" {
			value, err = numbervalidator.Validation(key, value, keyrules)
		} else if keyrule["Type"] == "Float" {
			value, err = floatvalidator.Validation(key, value, keyrules)
		} else if keyrule["Type"] == "DateTime" {
			value, err = datetimevalidator.Validation(key, value, keyrules)
		} else {
			value, err = stringvalidator.Validation(key, value, keyrules)
		}
		if err.httpcode != 200 || err.responsecode != 200 {
			return nil, err
		} else {
			keymap[key] = value
		}
	}
	return keymap, err
}
