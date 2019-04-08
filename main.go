package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", chrisHandler)
	http.ListenAndServe(":8080", nil)
}

func chrisHandler(w http.ResponseWriter, r *http.Request) {
	leagueStandings := getFantasyChris()

	leagueStandings.populateScores(getMLB())
	leagueStandings.rank()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	render(leagueStandings, w)
}
