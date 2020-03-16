package serviceerrors

type Error struct {
	httpcode        int
	responsecode    int
	errortype       string
	errorkey        string
	responsemessage string
}

func Okresponse() Error {
	return Error{
		httpcode:        200,
		responsecode:    200,
		errortype:       "",
		errorkey:        "",
		responsemessage: ""}
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
		httpcode:        200,
		responsecode:    400,
		errortype:       "INVALID_PARAM",
		errorkey:        param_name,
		responsemessage: message}
}
func ParamMissing(param_name string, errorsubtype string) Error {
	var message string
	if errorsubtype == "PARAMMISSING" {
		message = param_name + "'s missing"
	}
	return Error{
		httpcode:        200,
		responsecode:    400,
		errortype:       "PARAM_MISSING",
		errorkey:        param_name,
		responsemessage: message}
}

func InternalError(param_name string, errorsubtype string) Error {
	var message string
	if errorsubtype == "BADCONFIG" {
		message = "Internal Server Error"
	}
	return Error{
		httpcode:        200,
		responsecode:    500,
		errortype:       "INVALID_CONFIG",
		errorkey:        param_name,
		responsemessage: message}
}
