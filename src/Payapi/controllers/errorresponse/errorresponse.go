package errorresponse

import (
	"fmt"
	"net/http"
)

func NotfoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is error handler for path"+r.URL.Path[1:])
}
