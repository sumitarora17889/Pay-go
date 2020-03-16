package stringvalidator

import (
	"regexp"
	"serviceerrors"
	"strconv"
	"strings"
	"utils"
)

var Regex = map[string]string{"ALPHACAPS": "A-Z", //[A-Z]
	"ALPHANUM":   "A-Za-z0-9",      //[A-Za-z0-9]
	"ALPHASMALL": "a-z",            //[a-z]
	"ALPHAALL":   "A-Za-z",         //[A-Za-z]
	"ALLSAFE":    "A-Za-z0-9 .-_@", //
	"ALLCHARS":   "."}              //No Check

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
	value, error := isvalid(key, value, rules["Validationaction"], rules["Pattern"])
	if error.responsecode != 200 {
		return "", error
	}
	value, error = checklength(key, value, rules["Length"], rules["Validationaction"])
	if error.responsecode != 200 {
		return "", error
	}
	return value, serviceerrors.Okresponse()
}
func isvalid(key string, value string, Validationaction string, Pattern string) (string, serviceerrors.Error) {
	regex, found := Regex[Pattern]
	if found == false {
		return "", serviceerrors.InternalError(key, "BADCONFIG")
	} else {
		re, err := regexp.Compile("[" + regex + "]*")
		if err != nil {
			return "", serviceerrors.InternalError(key, "BADCONFIG")
		}
		if !re.MatchString(value) {
			if Validationaction == "blank" {
				return "", serviceerrors.Okresponse()
			} else if Validationaction == "trim" {
				re1, err1 := regexp.Compile("[^" + regex + "]*")
				if err1 != nil {
					return "", serviceerrors.InternalError(key, "BADCONFIG")
				}
				return re1.ReplaceAllString(value, ""), serviceerrors.Okresponse()
			} else {
				return "", serviceerrors.InvalidParams(key, "TYPEMISMATCH")
			}
		} else {
			return value, serviceerrors.Okresponse()
		}
	}
}

func checklength(key string, value string, length string, Validationaction string) (string, serviceerrors.Error) {
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
