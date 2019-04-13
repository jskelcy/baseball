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

type fantasypctLeague interface {
	sort.Interface
	Teams() []*fantasyTeam
	Rank()
}

type pctLeague struct {
	teams []*fantasyTeam
}

func (l *pctLeague) Teams() []*fantasyTeam {
	return l.teams
}

func (l *pctLeague) Len() int {
	return len(l.teams)
}

func (l *pctLeague) Less(i, j int) bool {
	if l.teams[i].perc > l.teams[j].perc {
		return true
	}
	return false
}

func (l *pctLeague) Swap(i, j int) {
	l.teams[i], l.teams[j] = l.teams[j], l.teams[i]
}

func (l *pctLeague) Rank() {
	sort.Sort(l)
	var currRank int
	var prevPerc float64

	for i, team := range l.teams {
		if prevPerc != team.perc {
			currRank = i + 1
		}
		team.Rank = currRank
		prevPerc = team.perc
	}
}

type winLeague struct {
	teams []*fantasyTeam
}

func (l *winLeague) Teams() []*fantasyTeam {
	return l.teams
}

func (l *winLeague) Len() int {
	return len(l.teams)
}

func (l *winLeague) Less(i, j int) bool {
	if l.teams[i].Wins > l.teams[j].Wins {
		return true
	}
	return false
}

func (l *winLeague) Swap(i, j int) {
	l.teams[i], l.teams[j] = l.teams[j], l.teams[i]
}

func (l *winLeague) Rank() {
	sort.Sort(l)
	var currRank int
	var prevWin int64

	for i, team := range l.teams {
		if prevWin != team.Wins {
			currRank = i + 1
		}
		team.Rank = currRank
		prevWin = team.Wins
	}
}

func populateScores(l fantasypctLeague, mlbScores mlbStandings) {
	for _, f := range l.Teams() {
		for _, t := range f.teams {
			f.Wins = f.Wins + mlbScores.Standing[t].Won
			f.Losses = f.Losses + mlbScores.Standing[t].Lost
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
