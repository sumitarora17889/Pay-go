package creditapplicationmodel

import "Payapi/components/serviceerrors"

func Createapplication(params map[string]interface{}) (map[string]interface{}, serviceerrors.Error) {
	return params, serviceerrors.Okresponse()
}
