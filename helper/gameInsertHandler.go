package helper

import (
	"net/http"
	dbfunc "ratingtable/1databasefunctionality"
)

func GameInsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		GameInsertExecution(w, r, dbfunc.Db)

	}
}
