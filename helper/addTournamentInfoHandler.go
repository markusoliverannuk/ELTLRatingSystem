package helper

import (
	"database/sql"
	"net/http"
)

var Db2 *sql.DB

func AddTournamentInfoHandler(w http.ResponseWriter, r *http.Request) {

	AddTournamentInfoWhenSubmitted(w, r, Db)
}
