package helper

import (
	"fmt"
	"net/http"
	types "ratingtable/1types"
)

func LoadRatingPage(w http.ResponseWriter, r *http.Request, players []types.Person) {
	data := struct {
		Players []types.Person
	}{
		Players: players,
	}

	e := ratingPageTMPLT.Execute(w, data)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "Error while rendering ELTL reiting!", http.StatusInternalServerError)
	}

}
