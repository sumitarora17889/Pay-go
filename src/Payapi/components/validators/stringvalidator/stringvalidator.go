package stringvalidator

import (
	"Payapi/components/serviceerrors"
	"regexp"
	"strconv"
	"strings"
)

var Regex = map[string]string{
	"ALPHACAPS":  "A-Z",            //[A-Z]
	"ALPHANUM":   "A-Za-z0-9",      //[A-Za-z0-9]
	"ALPHASMALL": "a-z",            //[a-z]
	"ALPHAALL":   "A-Za-z",         //[A-Za-z]
	"ALLSAFE":    "A-Za-z0-9 .-_@", //
	"ALLCHARS":   ".",
	"NUM":        "0-9"} //No Check

func Validation(key string, values string, rules map[string]string) (interface{}, serviceerrors.Error) {
	var val interface{}
	var finalval string
	var valarr []string
	var err serviceerrors.Error
	if rules["List"] == "1" {
		for _, value := range strings.Split(values, ",") {
			val, err = singlevalidation(key, value, rules)
			if err.Httpcode == 200 && err.Responsecode == 200 && val != nil {
				valarr = append(valarr, val.(string))
			} else {
				return nil, err
			}
		}
		finalval = strings.Join(valarr, ",")
	} else {
		val, err = singlevalidation(key, values, rules)
		return val, err
	}
	return finalval, err

}
func singlevalidation(key string, value string, rules map[string]string) (interface{}, serviceerrors.Error) {
	val, err := isvalid(key, value, rules["Validationaction"], rules["Pattern"])
	if err.Responsecode != 200 || val == nil {
		return nil, err
	}
	val, err = checklength(key, val.(string), rules["Length"], rules["Validationaction"])
	if err.Responsecode != 200 || val == nil {
		return nil, err
	}
	return value, serviceerrors.Okresponse()
}
func isvalid(key string, value string, Validationaction string, Pattern string) (interface{}, serviceerrors.Error) {
	regex, found := Regex[Pattern]
	if found == false {
		return nil, serviceerrors.InternalError(key, "BADCONFIG")
	} else {
		re, err := regexp.Compile("[" + regex + "]*")
		if err != nil {
			return nil, serviceerrors.InternalError(key, "BADCONFIG")
		}
		if !re.MatchString(value) {
			if Validationaction == "blank" {
				return nil, serviceerrors.Okresponse()
			} else if Validationaction == "trim" {
				re1, err1 := regexp.Compile("[^" + regex + "]*")
				if err1 != nil {
					return "", serviceerrors.InternalError(key, "BADCONFIG")
				}
				return re1.ReplaceAllString(value, ""), serviceerrors.Okresponse()
			} else {
				return nil, serviceerrors.InvalidParams(key, "TYPEMISMATCH")
			}
		} else {
			return value, serviceerrors.Okresponse()
		}
	}
}

func checklength(key string, value string, length string, Validationaction string) (interface{}, serviceerrors.Error) {
	leng, err := strconv.Atoi(length)
	if err != nil {
		return value, serviceerrors.Okresponse()
	} else {
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
