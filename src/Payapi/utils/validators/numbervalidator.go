package numbervalidator

import (
	"serviceerrors"
	"strconv"
	"strings"
)

func Validation(key string, values string, rules map[string]string) (interface{}, serviceerrors.Error) {
	var finalval []interface{}
	var err serviceerrors.Error
	if rules["List"] == "1" {
		for _, value := range strings.Split(values, ",") {
			val, err := singlevalidation(key, value, rules)
			if err.httpcode == 200 && err.responsecode == 200 && val != nil {
				finalval = append(finalval, val)
			} else {
				return nil, err
			}
		}
	} else {
		return singlevalidation(key, values, rules)
	}
	return finalval, err

}
func singlevalidation(key string, value string, rules map[string]string) (interface{}, serviceerrors.Error) {
	output, error := isvalid(key, value, rules["Validationaction"])
	if error.responsecode != 200 {
		return nil, error
	}
	output, error = checkrange(key, value, rules["Max"], rules["Min"], rules["Validationaction"])
	if error.responsecode != 200 {
		return nil, error
	}
	output, error = checklength(key, value, rules["Length"], rules["Validationaction"])
	if error.responsecode != 200 {
		return nil, error
	}
	return output, serviceerrors.Okresponse()
}
func isvalid(key string, value string, Validationaction string) (interface{}, serviceerrors.Error) {
	output, err := strconv.Atoi(value)
	if err != nil {
		if Validationaction == "blank" {
			return nil, serviceerrors.Okresponse()
		} else {
			return nil, serviceerrors.InvalidParams(key, "TYPEMISMATCH")
		}
	} else {
		return output, serviceerrors.Okresponse()
	}
}

func checkrange(key string, value string, max string, min string, Validationaction string) (interface{}, serviceerrors.Error) {
	a, _ := strconv.Atoi(value)
	rulesmin, errmin := strconv.Atoi(min)
	rulesmax, errmax := strconv.Atoi(max)
	if (errmin != nil || a >= rulesmin) && (errmax != nil || a <= rulesmax) {
		return value, serviceerrors.Okresponse()
	} else {
		if Validationaction == "blank" {
			return nil, serviceerrors.Okresponse()
		} else {
			return nil, serviceerrors.InvalidParams(key, "OUTOFRANGE")
		}
	}
}
func checklength(key string, value string, length string, Validationaction string) (interface{}, serviceerrors.Error) {
	leng, err := strconv.Atoi(length)
	a, _ := strconv.Atoi(value)
	if err != nil {
		return value, serviceerrors.Okresponse()
	} else {
		if len(value) <= leng {
			return a, serviceerrors.Okresponse()
		} else {
			if Validationaction == "blank" {
				return nil, serviceerrors.Okresponse()
			} else {
				return nil, serviceerrors.InvalidParams(key, "LENGTHEXCEEDED")
			}
		}
	}
}
