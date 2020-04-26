package serviceerrors

type Error struct {
	Httpcode        int
	Responsecode    int
	Errortype       string
	Errorkey        string
	Responsemessage string
}

func Okresponse() Error {
	return Error{
		Httpcode:        200,
		Responsecode:    200,
		Errortype:       "",
		Errorkey:        "",
		Responsemessage: ""}
}

func InvalidParams(param_name string, errorsubtype string) Error {
	var message string
	if errorsubtype == "LENGTHEXCEEDED" {
		message = "Bad Request. " + param_name + "'s value outside accepted range"
	} else if errorsubtype == "OUTOFRANGE" {
		message = "Bad Request. " + param_name + "'s value outside accepted range"
	} else if errorsubtype == "TYPEMISMATCH" {
		message = "Bad Request. Invalid " + param_name
	}
	return Error{
		Httpcode:        200,
		Responsecode:    400,
		Errortype:       "INVALID_PARAM",
		Errorkey:        param_name,
		Responsemessage: message}
}
func ParamMissing(param_name string, errorsubtype string) Error {
	var message string
	if errorsubtype == "PARAMMISSING" {
		message = param_name + "'s missing"
	}
	return Error{
		Httpcode:        200,
		Responsecode:    400,
		Errortype:       "PARAM_MISSING",
		Errorkey:        param_name,
		Responsemessage: message}
}

func InternalError(param_name string, errorsubtype string) Error {
	var message string
	if errorsubtype == "BADCONFIG" {
		message = "Internal Server Error"
	}
	return Error{
		Httpcode:        200,
		Responsecode:    500,
		Errortype:       "INVALID_CONFIG",
		Errorkey:        param_name,
		Responsemessage: message}
}
