package types

// Players in leaderboard ---------------------------

type Person struct {
	RankingPos int
	PersonID   int    `xml:"personid"`
	FamName    string `xml:"famname"`
	FirstName  string `xml:"firstname"`
	Sex        string `xml:"sex"`
	BirthDate  string `xml:"birthdate"`
	RateDate   string `xml:"ratedate"`
	RateOrder  int    `xml:"rateorder"`
	RatePlpnts int    `xml:"rateplpnts"`
	RatePoints int    `xml:"ratepoints"`
	RateWeight int    `xml:"rateweight"`
	RateWeilow string `xml:"rateweilow"`
	ClbName    string `xml:"clbname"`
}

type Persons struct {
	People []Person `xml:"person_x_reiting"`
}

// Tournaments in leaderboard -------------------------
type Tournament struct {
	AddedWhen       int
	Name            string
	HappenedWhen    string
	AmountOfPlayers string
	MainJudge       string
}

type Tournaments struct {
	TournamentList []Tournament
}
