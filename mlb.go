package main

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type mlbTeam struct {
	AwayLost                 int     `json:"away_lost"`
	AwayWon                  int     `json:"away_won"`
	Conference               string  `json:"conference"`
	ConferenceLost           int     `json:"conference_lost"`
	ConferenceWon            int     `json:"conference_won"`
	Division                 string  `json:"division"`
	DivisionLost             int     `json:"division_lost"`
	DivisionWon              int     `json:"division_won"`
	FirstName                string  `json:"first_name"`
	GamesBack                float64 `json:"games_back"`
	GamesPlayed              int     `json:"games_played"`
	HomeLost                 int     `json:"home_lost"`
	HomeWon                  int     `json:"home_won"`
	LastFive                 string  `json:"last_five"`
	LastName                 string  `json:"last_name"`
	LastTen                  string  `json:"last_ten"`
	Lost                     int64   `json:"lost"`
	OrdinalRank              string  `json:"ordinal_rank"`
	PointDifferential        int     `json:"point_differential"`
	PointDifferentialPerGame string  `json:"point_differential_per_game"`
	PointsAgainst            int     `json:"points_against"`
	PointsAllowedPerGame     string  `json:"points_allowed_per_game"`
	PointsFor                int     `json:"points_for"`
	PointsScoredPerGame      string  `json:"points_scored_per_game"`
	Rank                     int     `json:"rank"`
	Streak                   string  `json:"streak"`
	StreakTotal              int     `json:"streak_total"`
	StreakType               string  `json:"streak_type"`
	TeamID                   string  `json:"team_id"`
	WinPercentage            string  `json:"win_percentage"`
	Won                      int64   `json:"won"`
}

type mlbStandings struct {
	StandingsDate string
	Standing      map[string]mlbTeam
}

type mlbAPIStandings struct {
	StandingsDate string    `json:"standings_date"`
	Standing      []mlbTeam `json:"standing"`
}
type mlbClient struct {
	token         string
	userAgent     string
	currStandings mlbStandings
}

func (m *mlbClient) getMLBStandings() mlbStandings {
	return m.currStandings
}

func (m *mlbClient) Init() {
	if err := m.refreshCache(); err != nil {
		fmt.Printf("error initializing cache (will retry) %v", err)
	}
	fmt.Println("cache initialized")
	t := time.NewTicker(time.Minute * 5)

	go func() {
		for {
			<-t.C
			if err := m.refreshCache(); err != nil {
				fmt.Printf("error refreshing cache %v", err)
			}
			fmt.Println("updated cache")
		}
	}()
}

func (m *mlbClient) refreshCache() error {
	out := mlbStandings{
		Standing: make(map[string]mlbTeam),
	}
	rawStandings, err := compressedCall(m.token, m.userAgent)
	if err != nil {
		return err
	}

	for _, team := range rawStandings.Standing {
		out.Standing[team.TeamID] = team
	}

	m.currStandings = out
	return nil
}

func compressedCall(token, UserAgent string) (mlbAPIStandings, error) {
	mlbURL, _ := url.Parse("https://erikberg.com/mlb/standings.json")
	req := &http.Request{
		Method: http.MethodGet,
		URL:    mlbURL,
		Header: http.Header{
			"Accept-Encoding": []string{"gzip"},
			"Authorization":   []string{fmt.Sprintf("Bearer %s", token)},
			"User-Agent":      []string{UserAgent},
		},
		Close: true,
	}

	out := mlbAPIStandings{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return out, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return out, errors.New("cool your jets for a moment and try again")
	}
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return out, err
	}
	decoder := json.NewDecoder(gz)
	err = decoder.Decode(&out)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return out, err
	}

	return out, nil
}
