package main

import (
	"fmt"
	"sort"
)

type fantasyTeam struct {
	teams []string
	perc  float64

	Name   string    `json:"name"`
	Owner  string    `json:"owner"`
	Wins   int64     `json:"wins"`
	Losses int64     `json:"losses"`
	Perc   string    `json:"Perc"`
	Rank   int       `json:"rank"`
	Teams  []mlbTeam `json:"Teams"`
}

type fantasypctLeague interface {
	sort.Interface
	GetTeams() []*fantasyTeam
	Rank()
}

type pctLeague struct {
	Teams []*fantasyTeam `json:"teams"`
}

func (l *pctLeague) GetTeams() []*fantasyTeam {
	return l.Teams
}

func (l *pctLeague) Len() int {
	return len(l.Teams)
}

func (l *pctLeague) Less(i, j int) bool {
	if l.Teams[i].perc > l.Teams[j].perc {
		return true
	}
	return false
}

func (l *pctLeague) Swap(i, j int) {
	l.Teams[i], l.Teams[j] = l.Teams[j], l.Teams[i]
}

func (l *pctLeague) Rank() {
	sort.Sort(l)
	var currRank int
	var prevPerc float64

	for i, team := range l.Teams {
		if prevPerc != team.perc {
			currRank = i + 1
		}
		team.Rank = currRank
		prevPerc = team.perc
	}
}

type winLeague struct {
	Teams []*fantasyTeam `json:"teams"`
}

func (l *winLeague) GetTeams() []*fantasyTeam {
	return l.Teams
}

func (l *winLeague) Len() int {
	return len(l.Teams)
}

func (l *winLeague) Less(i, j int) bool {
	if l.Teams[i].Wins > l.Teams[j].Wins {
		return true
	}
	return false
}

func (l *winLeague) Swap(i, j int) {
	l.Teams[i], l.Teams[j] = l.Teams[j], l.Teams[i]
}

func (l *winLeague) Rank() {
	sort.Sort(l)
	var currRank int
	var prevWin int64

	for i, team := range l.Teams {
		if prevWin != team.Wins {
			currRank = i + 1
		}
		team.Rank = currRank
		prevWin = team.Wins
	}
}

func populateScores(l fantasypctLeague, mlbScores mlbStandings) {
	for _, f := range l.GetTeams() {
		for _, t := range f.teams {
			f.Wins = f.Wins + mlbScores.Standing[t].Won
			f.Losses = f.Losses + mlbScores.Standing[t].Lost
			if f.Losses != 0 {
				f.perc = float64(f.Wins) / float64(f.Losses+f.Wins)
			} else {
				f.perc = 1
			}
			f.Teams = append(f.Teams, mlbScores.Standing[t])
			f.Perc = fmt.Sprintf("%.3f", f.perc)
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
