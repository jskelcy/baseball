package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	s := server{
		apiKey:    os.Getenv("TOKEN"),
		userAgent: os.Getenv("USER_AGENT"),
	}
	fmt.Println(s.apiKey)
	fmt.Println(s.userAgent)

	http.HandleFunc("/", s.chrisHandler)
	http.HandleFunc("/CORK", s.ryanHandler)
	http.ListenAndServe(":8081", nil)
}

type server struct {
	apiKey    string
	userAgent string
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
}

func sendErr(w http.ResponseWriter, err error) {
	respBody := responseBody{
		Status: http.StatusInternalServerError,
		Error:  err.Error(),
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(respBody); err != nil {
		fmt.Println(err)
	}
}

type responseBody struct {
	Status int            `json:"status"`
	Error  string         `json:"error"`
	Teams  []*fantasyTeam `json:"teams"`
}

func (s *server) chrisHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	leagueStandings := getFantasyChris()
	mlb, err := getMLBAPI(s.apiKey, s.userAgent)
	if err != nil {
		sendErr(w, err)
		return
	}

	populateScores(leagueStandings, mlb)
	leagueStandings.Rank()

	respBody := responseBody{
		Status: http.StatusOK,
		Teams:  leagueStandings.Teams,
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(respBody); err != nil {
		fmt.Println(err)
	}
}

func (s *server) ryanHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	leagueStandings := getFantasyRyan()
	mlb, err := getMLBAPI(s.apiKey, s.userAgent)
	if err != nil {
		sendErr(w, err)
		return
	}

	populateScores(leagueStandings, mlb)
	leagueStandings.Rank()

	respBody := responseBody{
		Status: http.StatusOK,
		Teams:  leagueStandings.Teams,
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(respBody); err != nil {
		fmt.Println(err)
	}
}
