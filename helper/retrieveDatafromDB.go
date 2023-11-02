package helper

import (
	"fmt"
	"log"
)

func RetrievePlayerDataFromDatabase() ([]Person, error) {

	rows, err := Db.Query("SELECT * FROM leaderboard ORDER BY PP DESC")
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var players []Person
	rank := 1
	for rows.Next() {

		var player Person
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

func RetrieveTournamentsDataFromDatabase() ([]Tournament, error) {
	var tournaments []Tournament

	rows, err := Db.Query("SELECT * FROM tournaments")
	if err != nil {
		log.Fatal(err)
		return tournaments, err
	}
	defer rows.Close()

	for rows.Next() {
		var tournament Tournament
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
