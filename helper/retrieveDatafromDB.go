package helper

import (
	"fmt"
	"log"
	dbfunc "ratingtable/1databasefunctionality"
	types "ratingtable/1types"
)

func RetrievePlayerDataFromDatabase() ([]types.Person, error) {

	rows, err := dbfunc.Db.Query("SELECT * FROM leaderboard ORDER BY PP DESC")
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var players []types.Person
	rank := 1
	for rows.Next() {

		var player types.Person
		player.RankingPos = rank
		err := rows.Scan(
			&player.RatePlpnts,
			&player.RatePoints,
			&player.RateWeight,
			&player.PersonID,
			&player.FirstName,
			&player.FamName,
			&player.BirthDate,
			&player.ClbName,
		)
		if err != nil {
			// Handle the error
			fmt.Println(err, "heyy")

		}

		players = append(players, player)
		rank++
	}

	return players, nil
}

func RetrieveTournamentsDataFromDatabase() ([]types.Tournament, error) {
	var tournaments []types.Tournament

	rows, err := dbfunc.Db.Query("SELECT * FROM tournaments")
	if err != nil {
		log.Fatal(err)
		return tournaments, err
	}
	defer rows.Close()

	for rows.Next() {
		var tournament types.Tournament
		err := rows.Scan(
			&tournament.AddedWhen,
			&tournament.Name,
			&tournament.HappenedWhen,
			&tournament.AmountOfPlayers,
			&tournament.MainJudge,
		)
		if err != nil {
			log.Fatal(err)
			return tournaments, err
		}
		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}
