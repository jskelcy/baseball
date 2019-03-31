package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/baseball", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mlbStandings := getTeams()
	leagueStandings := getFantasy()
	for _, f := range *leagueStandings {
		for _, t := range f.teams {
			f.Wins = f.Wins + mlbStandings[t].wins
			f.Losses = f.Losses + mlbStandings[t].losses
			if f.Losses != 0 {
				f.Perc = fmt.Sprintf("%.3f", (float64(f.Wins) / float64(f.Losses+f.Wins)))
			} else {
				f.Perc = "1"
			}
		}
	}

	leagueStandings.rank()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	render(leagueStandings, w)
}
