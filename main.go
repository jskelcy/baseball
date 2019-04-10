package main

import (
	"flag"
	"net/http"
)

func main() {
	tokenPtr := flag.String("token", "foo", "a string")
	s := server{
		apiKey: *tokenPtr,
	}
	http.HandleFunc("/baseball", s.chrisHandler)
	http.ListenAndServe(":8080", nil)
}

type server struct {
	apiKey string
}

func (s *server) chrisHandler(w http.ResponseWriter, r *http.Request) {
	leagueStandings := getFantasyChris()

	leagueStandings.populateScores(getMLBAPI(s.apiKey))
	leagueStandings.rank()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	render(leagueStandings, w)
}
