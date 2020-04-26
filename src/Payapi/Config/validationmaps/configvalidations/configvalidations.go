package configvalidations

/*
	Length: The length of the variable
	Precision: The places after decimal the variable can go
	Type: Number/Char/Float/Boolean/Auth
	List: Whether the variable is array or single
	Required: Whether the variable is critical for the service or not.
	Validationaction: Blank/Trim/error
	Max: maximum value a number can attain
	Min: minimum value a number can attain
*/

var CreateApplicationMap map[string]map[string]string = map[string]map[string]string{
	"credit_id":             map[string]string{"Length": "10", "Precision": "0", "Type": "Number", "Pattern": "", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"application_id":        map[string]string{"Length": "10", "Precision": "0", "Type": "Number", "Pattern": "", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_id":              map[string]string{"Length": "10", "Precision": "0", "Type": "Number", "Pattern": "", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_first_name":      map[string]string{"Length": "30", "Precision": "0", "Type": "String", "Pattern": "ALLSAFE", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_last_name":       map[string]string{"Length": "30", "Precision": "0", "Type": "String", "Pattern": "ALLSAFE", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_email":           map[string]string{"Length": "100", "Precision": "0", "Type": "String", "Pattern": "ALLSAFE", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_mobile":          map[string]string{"Length": "10", "Precision": "0", "Type": "String", "Pattern": "NUM", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_addr1":           map[string]string{"Length": "100", "Precision": "0", "Type": "String", "Pattern": "ALLSAFE", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "trim"},
	"payer_addr2":           map[string]string{"Length": "100", "Precision": "0", "Type": "String", "Pattern": "ALLSAFE", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "trim"},
	"payer_pincode":         map[string]string{"Length": "6", "Precision": "0", "Type": "Number", "Pattern": "NUM", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_dob":             map[string]string{"Length": "10", "Precision": "0", "Type": "DateTime", "Pattern": "", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_pan":             map[string]string{"Length": "10", "Precision": "0", "Type": "String", "Pattern": "ALPHANUM", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_gst":             map[string]string{"Length": "15", "Precision": "0", "Type": "String", "Pattern": "ALPHANUM", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_aadhaar":         map[string]string{"Length": "12", "Precision": "0", "Type": "String", "Pattern": "NUM", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"pan_file":              map[string]string{"Length": "600", "Precision": "0", "Type": "String", "Pattern": "ALLCHARS", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"aadhaar_file":          map[string]string{"Length": "600", "Precision": "0", "Type": "String", "Pattern": "ALLCHARS", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"gst_file":              map[string]string{"Length": "600", "Precision": "0", "Type": "String", "Pattern": "ALLCHARS", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"selfie_file":           map[string]string{"Length": "600", "Precision": "0", "Type": "String", "Pattern": "ALLCHARS", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_bank_name":       map[string]string{"Length": "100", "Precision": "0", "Type": "String", "Pattern": "ALLSAFE", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_bank_acc_no":     map[string]string{"Length": "15", "Precision": "0", "Type": "String", "Pattern": "ALPHANUM", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_acc_name":        map[string]string{"Length": "100", "Precision": "0", "Type": "String", "Pattern": "ALPHAALL", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_bank_ifsc":       map[string]string{"Length": "11", "Precision": "0", "Type": "String", "Pattern": "ALPHANUM", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"payer_acc_type":        map[string]string{"Length": "1", "Precision": "0", "Type": "String", "Pattern": "ALPHACAPS", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "error"},
	"success_redirect_url":  map[string]string{"Length": "600", "Precision": "0", "Type": "String", "Pattern": "ALLCHARS", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "blank"},
	"failure_redirect_url":  map[string]string{"Length": "600", "Precision": "0", "Type": "String", "Pattern": "ALLCHARS", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "blank"},
	"tnc_accepted":          map[string]string{"Length": "1", "Precision": "0", "Type": "String", "Pattern": "ALPHACAPS", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"platform":              map[string]string{"Length": "10", "Precision": "0", "Type": "String", "Pattern": "ALPHACAPS", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"user_modid":            map[string]string{"Length": "10", "Precision": "0", "Type": "String", "Pattern": "ALPHACAPS", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"user_current_url":      map[string]string{"Length": "600", "Precision": "0", "Type": "String", "Pattern": "ALLCHARS", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "trim"},
	"user_longitude":        map[string]string{"Length": "12", "Precision": "8", "Type": "Float", "Pattern": "", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "blank"},
	"user_latitude":         map[string]string{"Length": "12", "Precision": "8", "Type": "Float", "Pattern": "", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "blank"},
	"user_gaid":             map[string]string{"Length": "50", "Precision": "0", "Type": "String", "Pattern": "ALLSAFE", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "blank"},
	"user_ip_iso":           map[string]string{"Length": "2", "Precision": "0", "Type": "String", "Pattern": "ALPHACAPS", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"user_ip":               map[string]string{"Length": "15", "Precision": "0", "Type": "String", "Pattern": "ALLSAFE", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"user_ip_city_id":       map[string]string{"Length": "10", "Precision": "0", "Type": "Number", "Pattern": "", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "blank"},
	"user_ip_accuracy":      map[string]string{"Length": "4", "Precision": "0", "Type": "Number", "Pattern": "", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "blank"},
	"user_login_mode":       map[string]string{"Length": "1", "Precision": "0", "Type": "Number", "Pattern": "", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "blank"},
	"user_agent":            map[string]string{"Length": "320", "Precision": "0", "Type": "String", "Pattern": "ALPHAALL", "List": "0", "Required": "False", "Max": "", "Min": "", "Validationaction": "blank"},
	"usrid":                 map[string]string{"Length": "10", "Precision": "0", "Type": "Number", "Pattern": "", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
	"user_current_pagename": map[string]string{"Length": "100", "Precision": "0", "Type": "String", "Pattern": "ALPHAALL", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "blank"},
	"timestamp":             map[string]string{"Length": "20", "Precision": "0", "Type": "DateTime", "Pattern": "", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "blank"},
	"user_auth_key":         map[string]string{"Length": "128", "Precision": "0", "Type": "String", "Pattern": "ALPHANUM", "List": "0", "Required": "True", "Max": "", "Min": "", "Validationaction": "error"},
}
