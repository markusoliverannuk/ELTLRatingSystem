package helper

import (
	"database/sql"
	"log"
)

func CreateDbTablePlayers(db *sql.DB) {
	_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS leaderboard (
		PP INTEGER,
        RP INTEGER,
        KL INTEGER,
        ID INTEGER,
        EESNIMI TEXT,
        PERENIMI TEXT,
        SA INTEGER,
        KLUBI TEXT
    )`)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func CreateDbTableTournaments(db *sql.DB) {
	_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS tournaments (
        LISATUD STRING,
        NIMI STRING,
        TOIMUMISKUUP2EV STRING,
        M2NGIJATEARV INT,
       PEAKOHTUNIK TEXT
       
    )`)
	if err1 != nil {
		log.Fatal(err1)
	}
}
