package main

func getFantasyRyan() *winLeague {
	return &winLeague{
		teams: []*fantasyTeam{
			{
				Name:  "Thudbutt",
				Owner: "Manso",
				teams: []string{BO, TB, SL, SD, CIN, MI},
			},
			{
				Name:  "deGromNomNom",
				Owner: "Kate",
				teams: []string{BO, ATL, OAK, TO, BA, KC},
			},
			{
				Name:  "John Krunk All-Stars",
				Owner: "Marc",
				teams: []string{HO, TB, SL, SD, DE, MI},
			},
			{
				Name:  "Ryan & Katie",
				Owner: "UltraViolet",
				teams: []string{HO, SE, CHIC, CO, SD, MI},
			},
			{
				Name:  "BlöndeHammer",
				Owner: "Amanda",
				teams: []string{NYY, WA, NYM, PITT, SF, KC},
			},
			{
				Name:  "Purple Reign",
				Owner: "Niccole",
				teams: []string{HO, WA, MIN, CO, TO, BA},
			},
			{
				Name:  "FU Jobu",
				Owner: "Walker",
				teams: []string{LAD, WA, ATL, LAA, PITT, KC},
			},
			{
				Name:  "Diaper Doody",
				Owner: "Billy",
				teams: []string{NYY, MIL, NYM, LAA, PITT, CHIW},
			},
			{
				Name:  "Wendy Preffercorn",
				Owner: "Bryce",
				teams: []string{CL, TB, NYM, OAK, CIN, TX},
			},
			{
				Name:  "JeterFrazierPhoto",
				Owner: "Chris",
				teams: []string{CL, MIL, ATL, OAK, AZ, SF},
			},
			{
				Name:  "ARod's Juicebox",
				Owner: "Paul",
				teams: []string{NYY, PHIL, SE, LAA, AZ, DE},
			},
			{
				Name:  "Jersey Jackass",
				Owner: "Matt",
				teams: []string{LAD, MIL, SE, MIN, DE, TX},
			},
			{
				Name:  "Whoopie Pies",
				Owner: "Pops",
				teams: []string{BO, PHIL, AZ, CIN, BA, TX},
			},
			{
				Name:  "Gutterballs",
				Owner: "Pete",
				teams: []string{LAD, SL, MIN, CHIC, SF, CHIW},
			},
			{
				Name:  "Bowman Bombers",
				Owner: "Russ",
				teams: []string{PHIL, CL, CHIC, CO, CHIW, TO},
			},
		},
	}
}
