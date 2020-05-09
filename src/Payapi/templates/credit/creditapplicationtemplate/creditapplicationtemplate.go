package creditapplicationtemplate

import "Payapi/components/serviceerrors"

func Createapplicationresponse(output map[string]interface{}) (interface{}, serviceerrors.Error) {
	return output, serviceerrors.Okresponse()
}
