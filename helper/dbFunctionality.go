package helper

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func DatabaseInitialization(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	return db
}

func InsertPlayerToLeaderboard(db *sql.DB, person Person) error {

	birthDate, err := time.Parse("2006-01-02", person.BirthDate)
	if err != nil {
		return err
	}

	birthYear := birthDate.Year()
	insertSQL := `
		        INSERT INTO leaderboard ( PP, RP, KL, ID, EESNIMI, PERENIMI, SA, KLUBI)
				VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)
		    `

	_, err3 := db.Exec(
		insertSQL,
		person.RatePlpnts,
		person.RatePoints,
		person.RateWeight,
		person.PersonID,
		person.FirstName,
		person.FamName,
		birthYear,
		person.ClbName,
	)

	fmt.Println(person)
	return err3

}

func ProcessXMLFile(filename string) error {
	var count int
	err := Db.QueryRow("SELECT COUNT(*) FROM leaderboard").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		xmlFile, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer xmlFile.Close()

		var persons Persons
		decoder := xml.NewDecoder(xmlFile)
		if err := decoder.Decode(&persons); err != nil {
			return err
		}

		for _, person := range persons.People {
			err := InsertPlayerToLeaderboard(Db, person)
			if err != nil {
				fmt.Printf("Error inserting player %d: %v\n", person.PersonID, err)
			}
		}
	}
	return nil
}
