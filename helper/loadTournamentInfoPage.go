package helper

import (
	"fmt"
	"net/http"
	types "ratingtable/1types"
)

func LoadTournamentAddPage(w http.ResponseWriter, r *http.Request, tournaments []types.Tournament) {
	data := struct {
		Tournaments []types.Tournament
	}{
		Tournaments: tournaments,
	}

	e := tournamentAddPageTMPLT.Execute(w, data)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "Error while rendering tournament adder page!", http.StatusInternalServerError)
	}
}
