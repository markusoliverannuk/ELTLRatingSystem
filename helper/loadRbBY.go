package helper

import (
	"log"
	dbfunc "ratingtable/1databasefunctionality"
	types "ratingtable/1types"
)

func RetrievePlayerDataFromDatabaseByBirthYear(birthyear string) ([]types.Person, error) {
	var players []types.Person

	rows, err := dbfunc.Db.Query("SELECT * FROM leaderboard WHERE SA LIKE ?", "%"+birthyear+"%")
	if err != nil {
		log.Fatal(err)
		return players, err
	}
	defer rows.Close()

	for rows.Next() {
		var player types.Person
		err := rows.Scan(
			&player.RateOrder,
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
			log.Fatal(err)
			return players, err
		}
		players = append(players, player)
	}

	return players, nil
}
