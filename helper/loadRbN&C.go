package helper

import (
	"fmt"
	"log"
	dbfunc "ratingtable/1databasefunctionality"
	types "ratingtable/1types"
)

func RetrievePlayerDataFromDatabaseByNameAndClub(playername string, clubname string) ([]types.Person, error) {
	var players []types.Person
	fmt.Println("Used")
	rows, err := dbfunc.Db.Query("SELECT * FROM leaderboard WHERE (EESNIMI LIKE ? OR PERENIMI LIKE ?) AND KLUBI LIKE ?", "%"+playername+"%", "%"+playername+"%", "%"+clubname+"%")
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
