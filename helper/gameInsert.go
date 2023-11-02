package helper

import (
	"database/sql"
	"fmt"
	"net/http"
)

func GameInsertExecution(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == "POST" {
		WinnerPlayer := r.FormValue("winnerplayer")

		LoserPlayer := r.FormValue("loserplayer")

		ratingPointsWinner, err := getRatingPLPoints(db, WinnerPlayer)
		if err != nil {
			fmt.Println("Error fetching rating points for winner:", err)
			return
		}

		ratingPointsLoser, err := getRatingPLPoints(db, LoserPlayer)
		if err != nil {
			fmt.Println("Error fetching rating points for loser:", err)
			return
		}

		weightPointsWinner, err := getWeightPoints(db, WinnerPlayer)
		if err != nil {
			fmt.Println("Error getting weight points for winner", err)
			return
		}
		weightPointsLoser, err := getWeightPoints(db, LoserPlayer)
		if err != nil {
			fmt.Println("Error getting weight points for loser:", err)
			return
		}
		fmt.Println(weightPointsWinner, " <---- Weight points for winner")
		fmt.Println(weightPointsLoser, " <---- Weight points for loser")

		var RPSubstraction int
		RPSubstraction = ratingPointsWinner - ratingPointsLoser
		if RPSubstraction < 0 {
			RPSubstraction *= -1
		}
		fmt.Println("Winner's rating Points:", ratingPointsWinner)
		fmt.Println("Loser's rating Points:", ratingPointsLoser)
		fmt.Println("Reitingupunktide vahe:", RPSubstraction)

		PointChangePreWeight := CalculateRPForPlayers(ratingPointsWinner, ratingPointsLoser, RPSubstraction)
		fmt.Println("Point Change before weights applied", PointChangePreWeight)
		FinalPointChangeWinner := AddWeightedMultiplier(PointChangePreWeight, weightPointsWinner)
		FinalPointChangeLoser := AddWeightedMultiplier(PointChangePreWeight, weightPointsLoser)
		FinalPointChangeLoser *= -1
		fmt.Println("Final point change with weightalgorithm added for winner:", FinalPointChangeWinner)
		fmt.Println("Final point change with weightalgorithm added for loser:", FinalPointChangeLoser)

		UpdateRanking(WinnerPlayer, FinalPointChangeWinner)
		UpdateRanking(LoserPlayer, FinalPointChangeLoser)


		http.Redirect(w, r, "/addtournament", http.StatusSeeOther)

		//Hommikul teha

		// Otsi databaseist nii winner playerit kui ka loser playerit, vaata kas on olemas, kui ei ole
		// ss lükka error, kui mõlemad on databaseis olemas siis otsi mõlema reitingupunktid ja võrdle neid
		// olenevalt sellest mis on punktide vahe, ss kas liida v lahuta kindel arv reitingupunkte neilt.
	}
}
