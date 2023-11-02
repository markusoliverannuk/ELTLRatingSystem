package helper

import (
	"fmt"
	dbfunc "ratingtable/1databasefunctionality"
)

func UpdateRanking(PlayerID string, PointAmount int) {
	var currentPP int
	var currentRP int

	err := dbfunc.Db.QueryRow("SELECT PP FROM leaderboard WHERE ID = ?", PlayerID).Scan(&currentPP)
	if err != nil {
		fmt.Println("encountered error finding PP:", err)
		return
	}
	err2 := dbfunc.Db.QueryRow("SELECT RP FROM leaderboard WHERE ID = ?", PlayerID).Scan(&currentRP)
	if err2 != nil {
		fmt.Println("encountered error finding RP:", err)
		return
	}

	updatedPP := currentPP + PointAmount
	updatedRP := currentRP + PointAmount

	_, updateErr := dbfunc.Db.Exec("UPDATE leaderboard SET PP = ? WHERE ID = ?", updatedPP, PlayerID)
	if updateErr != nil {
		fmt.Println("error updating PP:", updateErr)
		return
	}
	_, updateErr2 := dbfunc.Db.Exec("UPDATE leaderboard SET RP = ? WHERE ID = ?", updatedRP, PlayerID)
	if updateErr2 != nil {
		fmt.Println("error updating RP:", updateErr)
		return
	}

	fmt.Println("Updated ranking successfully.")
}
