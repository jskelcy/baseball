package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("TOKEN")
	userAgent := os.Getenv("USER_AGENT")
	fmt.Println(apiKey)
	fmt.Println(userAgent)

	client := &mlbClient{
		token:     apiKey,
		userAgent: userAgent,
	}
	client.Init()
	s := server{
		mlbClient: client,
	}

	fs := http.FileServer(http.Dir("./baseball_frontend/build/"))
	http.Handle("/", fs)
	http.HandleFunc("/CHRIS", s.chrisHandler)
	fmt.Println(http.ListenAndServe(":8081", nil))
}

type server struct {
	mlbClient *mlbClient
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
	mlb := s.mlbClient.getMLBStandings()

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
