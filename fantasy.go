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
	HO   = "Houston"
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
	PITT = "Pittsburgh"
	KC   = "Kansas City"
)
