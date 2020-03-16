package floatvalidator

import (
	"serviceerrors"
	"strconv"
	"strings"
	"utils"
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
	var error utils.Error
	if rules["List"] == "1" {
		for _, value := range strings.Split(values, ",") {
			val, error = singlevalidation(key, value, rules)
			if error.httpcode == 200 && error.responsecode == 200 && val != "" {
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
	if error.responsecode != 200 {
		return "", error
	}
	value, error = checkrange(key, value, rules["Max"], rules["Min"], rules["Validationaction"])
	if error.responsecode != 200 {
		return "", error
	}
	value, error = checklength(key, value, rules["Length"], rules["Validationaction"])
	if error.responsecode != 200 {
		return "", error
	}
	value, error = checkprecision(key, value, rules["Precision"], rules["Validationaction"])
	if error.responsecode != 200 {
		return "", error
	}
	return value, serviceerrors.Okresponse()
}
func isvalid(key string, value string, Validationaction string) (string, serviceerrors.Error) {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		if Validationaction == "blank" {
			return "", serviceerrors.Okresponse()
		} else {
			return "", serviceerrors.InvalidParams(key, "TYPEMISMATCH")
		}
	} else {
		return value, serviceerrors.Okresponse()
	}
}

func checkrange(key string, value string, max string, min string, Validationaction string) (string, serviceerrors.Error) {
	a, _ := strconv.ParseFloat(value, 64)
	rulesmin, errmin := strconv.ParseFloat(min, 64)
	rulesmax, errmax := strconv.ParseFloat(max, 64)
	if (errmin != nil || a >= rulesmin) && (errmax != nil || a <= rulesmax) {
		return value, serviceerrors.Okresponse()
	} else {
		if Validationaction == "blank" {
			return "", serviceerrors.Okresponse()
		} else {
			return "", serviceerrors.InvalidParams(key, "OUTOFRANGE")
		}
	}
}
func checklength(key string, value string, length string, Validationaction string) (string, serviceerrors.Error) {
	leng, err := strconv.Atoi(length)
	if err != nil {
		return value, serviceerrors.Okresponse()
	} else {
		if strings.Contains(value, ".") {
			leng = leng + 1
		}
		if len(value) <= leng {
			return value, serviceerrors.Okresponse()
		} else {
			if Validationaction == "blank" {
				return "", serviceerrors.Okresponse()
			} else if Validationaction == "trim" {
				return value[:leng], serviceerrors.Okresponse()
			} else {
				return "", serviceerrors.InvalidParams(key, "LENGTHEXCEEDED")
			}
		}
	}
}
func checkprecision(key string, value string, precision string, Validationaction string) (string, serviceerrors.Error) {
	prec, err := strconv.Atoi(precision)
	if err != nil {
		return value, serviceerrors.Okresponse()
	} else {
		if strings.Contains(value, ".") {
			val := strings.Split(value, '.')
			if len(val[1]) <= prec {
				return value, serviceerrors.Okresponse()
			} else {
				if Validationaction == "blank" {
					return "", serviceerrors.Okresponse()
				} else if Validationaction == "trim" {
					return val[0] + '.' + val[1][:prec], serviceerrors.Okresponse()
				} else {
					return "", serviceerrors.InvalidParams(key, "LENGTHEXCEEDED")
				}
			}
		} else {
			return value, serviceerrors.Okresponse()
		}
	}
}
