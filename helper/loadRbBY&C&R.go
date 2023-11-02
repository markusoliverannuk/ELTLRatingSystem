package helper

import (
	"fmt"
	"log"
	dbfunc "ratingtable/1databasefunctionality"
	types "ratingtable/1types"
	"strconv"
)

func RetrievePlayerDataFromDatabaseByBirthYearAndClubAndRange(birthyear string, clubname string, ratingrange string) ([]types.Person, error) {
	var players []types.Person

	ratingrangerune := []rune(ratingrange)

	fmt.Println(string(ratingrangerune[:3]))

	lowerRangePoint, err := strconv.Atoi((string(ratingrangerune[:3])))
	if err != nil {

	}
	upperRangePoint, err := strconv.Atoi((string(ratingrangerune[3:])))
	if err != nil {

	}
	rows, err := dbfunc.Db.Query("SELECT * FROM leaderboard WHERE SA LIKE ? AND KLUBI LIKE ? AND NR >= ? AND NR <= ?", "%"+birthyear+"%", "%"+clubname+"%", lowerRangePoint, upperRangePoint)
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
