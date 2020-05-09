package crediteligibility

import (
	"net/http"
)

var Actions map[string]func(http.ResponseWriter, *http.Request) = map[string]func(http.ResponseWriter, *http.Request){}
