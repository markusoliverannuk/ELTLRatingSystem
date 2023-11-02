package helper

import (
	"database/sql"
	"net/http"
	dbfunc "ratingtable/1databasefunctionality"
)

var Db2 *sql.DB

func AddTournamentInfoHandler(w http.ResponseWriter, r *http.Request) {

	AddTournamentInfoWhenSubmitted(w, r, dbfunc.Db)
}
