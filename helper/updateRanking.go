package helper

import "fmt"

func UpdateRanking(PlayerID string, PointAmount int) {
    var currentPP int
    err := Db.QueryRow("SELECT PP FROM leaderboard WHERE ID = ?", PlayerID).Scan(&currentPP)
    if err != nil {
        fmt.Println("encountered error finding PP:", err)
        return
    }

    updatedPP := currentPP + PointAmount

    _, updateErr := Db.Exec("UPDATE leaderboard SET PP = ? WHERE ID = ?", updatedPP, PlayerID)
    if updateErr != nil {
        fmt.Println("error updating PP:", updateErr)
        return
    }

    fmt.Println("Updated ranking successfully.")
}
