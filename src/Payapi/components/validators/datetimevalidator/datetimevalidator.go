package datetimevalidator

import (
	"Payapi/components/serviceerrors"
	"strings"
	"time"
)

func minimum(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Validation(key string, values string, rules map[string]string) (string, serviceerrors.Error) {
	var val string
	var finalval string
	var valarr []string
	var error serviceerrors.Error
	if rules["List"] == "1" {
		for _, value := range strings.Split(values, ",") {
			val, error = singlevalidation(key, value, rules)
			if error.Httpcode == 200 && error.Responsecode == 200 && val != "" {
				valarr = append(valarr, val)
			} else {
				return "", error
			}
		}
		finalval = strings.Join(valarr, ",")
	} else {
		finalval, error = singlevalidation(key, values, rules)
	}
	return finalval, error

}
func singlevalidation(key string, value string, rules map[string]string) (string, serviceerrors.Error) {
	value, error := isvalid(key, value, rules["Validationaction"])
	if error.Responsecode != 200 {
		return "", error
	}
	return value, serviceerrors.Okresponse()
}
func isvalid(key string, value string, Validationaction string) (string, serviceerrors.Error) {
	layout := "02-01-2006T15:04:05.000Z"
	datetime := strings.Split(value, "T")
	if len(datetime) != 2 {
		value = datetime[0] + "T" + "00:00:00.000Z"
	} else {
		value = datetime[0] + "T" + datetime[1]
	}
	t, err := time.Parse(layout, value)
	if err != nil {
		return "", serviceerrors.InvalidParams(key, "TYPEMISMATCH")
	} else {
		return t.String(), serviceerrors.Okresponse()
	}
}
