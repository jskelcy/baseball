package main

import "sort"

type fantasyTeam struct {
	Name   string
	Owner  string
	teams  []string
	Wins   int64
	Losses int64
	Perc   string
	Rank   int
}

type league []*fantasyTeam

func (l *league) Len() int {
	return len(*l)
}

func (l *league) Less(i, j int) bool {
	if (*l)[i].Wins > (*l)[j].Wins {
		return true
	}
	return false
}

func (l *league) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *league) rank() {
	sort.Sort(l)
	for i, team := range *l {
		team.Rank = i + 1
	}
}

// implent sort

const (
	NYY  = "N.Y. Yankees"
	NYM  = "N.Y. Mets"
	CHIC = "Chi. Cubs"
	CHIW = "Chi. White Sox"
	MI   = "Miami"
	WA   = "Washington"
	SL   = "St. Louis"
	LAA  = "L.A. Angels"
	LAD  = "L.A. Dodgers"
	CIN  = "Cincinnati"
	DE   = "Detroit"
	HO   = "Houstan"
	MIL  = "Milwaukee"
	OAK  = "Oakland"
	MIN  = "Minnesota"
	TX   = "Texas"
	PHIL = "Philadelphia"
	SD   = "San Diego"
	BO   = "Boston"
	SE   = "Seattle"
	CO   = "Colorado"
	BA   = "Baltimore"
	TO   = "Toronto"
	SF   = "San Francisco"
	CL   = "Cleveland"
	ATL  = "Atlanta"
	TB   = "Tampa Bay"
	AZ   = "Arizona"
	PITT = "Pittsburg"
	KC   = "Kansas City"
)

func getFantasy() *league {
	return &league{
		&fantasyTeam{
			Name:  "Bobby Bonilla’s $1.2 million",
			Owner: "Sean",
			teams: []string{NYY, WA, SL, LAA, CIN, DE},
		},
		&fantasyTeam{
			Name:  "50 Shades of Sonny Gray",
			Owner: "Harrison",
			teams: []string{HO, WA, MIL, OAK, CIN, DE},
		},
		&fantasyTeam{
			Name:  "Scrimblet ball",
			Owner: "Jake",
			teams: []string{NYY, WA, NYM, OAK, CHIW, MI},
		},
		&fantasyTeam{
			Name:  "Beffballs",
			Owner: "Beth",
			teams: []string{HO, NYM, MIL, MIN, CIN, TX},
		},
		&fantasyTeam{
			Name:  "DeGromNomNom",
			Owner: "Kate",
			teams: []string{NYY, PHIL, NYM, LAA, SD},
		},
		&fantasyTeam{
			Name:  "Game of Failure",
			Owner: "Matt",
			teams: []string{BO, WA, SL, OAK, CIN, SE},
		},
		&fantasyTeam{
			Name:  "Never Order Helmet Nachos",
			Owner: "Chuck",
			teams: []string{HO, PHIL, MIL, CO, SD, SE},
		},
		&fantasyTeam{
			Name:  "Luv2bsbl72",
			Owner: "Lindsay",
			teams: []string{BA, SL, TO, MIN, SF, DE},
		},
		&fantasyTeam{
			Name:  "a dream Chicken Finger Bucket",
			Owner: "Marie",
			teams: []string{NYM, CL, MIL, CO, CHIW, TX},
		},
		&fantasyTeam{
			Name:  "Relief Pitchers",
			Owner: "Chris",
			teams: []string{HO, CL, ATL, TB, SD, AZ},
		},
		&fantasyTeam{
			Name:  "Don't know nothin about baseball",
			Owner: "Keri",
			teams: []string{LAD, BA, SF, KC, TO, TX},
		},
		&fantasyTeam{
			Name:  "Jobu needs a refill",
			Owner: "Pat",
			teams: []string{CL, SL, TB, NYM, PITT, AZ},
		},
		&fantasyTeam{
			Name:  "No glove no love 2",
			Owner: "Katie",
			teams: []string{CHIC, SD, MIN, KC, BA, MI},
		},
		&fantasyTeam{
			Name:  "Ultraviolet",
			Owner: "Ryan",
			teams: []string{BO, PHIL, ATL, LAA, PITT, AZ},
		},
		&fantasyTeam{
			Name:  "Pedro Serrano",
			Owner: "Andrew",
			teams: []string{CL, CHIC, CO, ATL, CHIW, MI},
		},
		&fantasyTeam{
			Name:  "Murfreesboro Tongue Sandwiches",
			Owner: "Tyler",
			teams: []string{BO, TO, CHIW, BA, SE, KC},
		},
		&fantasyTeam{
			Name:  "I’m in a nightmare, empty chicken finger bucket",
			Owner: "Wolan",
			teams: []string{BO, CHIC, ATL, OAK, SF, PITT},
		},
		&fantasyTeam{
			Name:  "Zookey Burger",
			Owner: "Carol",
			teams: []string{LAD, PHIL, TB, CO, SF, SE},
		},
		&fantasyTeam{
			Name:  "wildxbaseball10",
			Owner: "Lauren",
			teams: []string{CHIC, LAD, MIN, TB, PITT, AZ},
		},
		&fantasyTeam{
			Name:  "HOME dRUNK",
			Owner: "Casey",
			teams: []string{LAA, LAD, KC, TO, DE, TX},
		},
	}
}
