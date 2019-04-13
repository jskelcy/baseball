package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	s := server{
		apiKey: os.Getenv("TOKEN"),
	}
	fmt.Println(s.apiKey)
	http.HandleFunc("/", s.chrisHandler)
	http.HandleFunc("/CORK", s.ryanHandler)
	http.ListenAndServe(":8080", nil)
}

type server struct {
	apiKey string
}

func (s *server) chrisHandler(w http.ResponseWriter, r *http.Request) {
	leagueStandings := getFantasyChris()
	mlb, err := getMLBAPI(s.apiKey)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	populateScores(leagueStandings, mlb)
	leagueStandings.Rank()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	render(leagueStandings, w)
}

func (s *server) ryanHandler(w http.ResponseWriter, r *http.Request) {
	leagueStandings := getFantasyRyan()
	mlb, err := getMLBAPI(s.apiKey)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	populateScores(leagueStandings, mlb)
	leagueStandings.Rank()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	render(leagueStandings, w)
}
