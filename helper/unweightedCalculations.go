package helper

func CalculateRPForPlayers(ratingPointsWinner int, ratingPointsLoser int, RPSubstraction int) int {

	var PointChange int
	if ratingPointsWinner >= ratingPointsLoser {
		if RPSubstraction >= 0 && RPSubstraction <= 2 {
			PointChange = 2
		} else if RPSubstraction > 2 && RPSubstraction <= 13 {
			PointChange = 1
		} else if RPSubstraction > 13 {
			PointChange = 0
		}
	} else if ratingPointsLoser > ratingPointsWinner {
		PointChange = (RPSubstraction + 5) / 3
	}

	return PointChange
}
