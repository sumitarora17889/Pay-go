package lending

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"database/sql"
)


func getappdetails(w http.ResponseWriter, r *http.Request)
{
	var err error
	var dbConn *sql.DB
	var start, end float64
	var inputParams, serviceResponse = make(map[string]interface{}), make(map[string]interface{})
	inputParams, err = utils.ReceiveRequestAndParse(r)
}