package main

import (
	"fmt"
	"sort"
)

type fantasyTeam struct {
	Name       string
	Owner      string
	teams      []string
	Wins       int64
	Losses     int64
	perc       float64
	RenderPerc string
	Rank       int
}

type league []*fantasyTeam

func (l *league) Len() int {
	return len(*l)
}

func (l *league) Less(i, j int) bool {
	if (*l)[i].perc > (*l)[j].perc {
		return true
	}
	return false
}

func (l *league) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *league) rank() {
	sort.Sort(l)
	var currRank int
	var prevPerc float64

	for i, team := range *l {
		if prevPerc != team.perc {
			currRank = i + 1
		}
		team.Rank = currRank
		prevPerc = team.perc
	}
}

func (l *league) populateScores(mlbScores mlb) {
	for _, f := range *l {
		for _, t := range f.teams {
			f.Wins = f.Wins + mlbScores[t].wins
			f.Losses = f.Losses + mlbScores[t].losses
			if f.Losses != 0 {
				f.perc = float64(f.Wins) / float64(f.Losses+f.Wins)
			} else {
				f.perc = 1
			}
			f.RenderPerc = fmt.Sprintf("%.3f", f.perc)
		}
	}
}

const (
	NYY  = "new-york-yankees"
	NYM  = "new-york-mets"
	CHIC = "chicago-cubs"
	CHIW = "chicago-white-sox"
	MI   = "miami-marlins"
	WA   = "washington-nationals"
	SL   = "st-louis-cardinals"
	LAA  = "los-angeles-angels"
	LAD  = "los-angeles-dodgers"
	CIN  = "cincinnati-reds"
	DE   = "detroit-tigers"
	HO   = "houston-astros"
	MIL  = "milwaukee-brewers"
	OAK  = "oakland-athletics"
	MIN  = "minnesota-twins"
	TX   = "texas-rangers"
	PHIL = "philadelphia-phillies"
	SD   = "san-diego-padres"
	BO   = "boston-red-sox"
	SE   = "seattle-mariners"
	CO   = "colorado-rockies"
	BA   = "baltimore-orioles"
	TO   = "toronto-blue-jays"
	SF   = "san-francisco-giants"
	CL   = "cleveland-indians"
	ATL  = "atlanta-braves"
	TB   = "tampa-bay-rays"
	AZ   = "arizona-diamondbacks"
	PITT = "pittsburgh-pirates"
	KC   = "kansas-city-royals"
)
