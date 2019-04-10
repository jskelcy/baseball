package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gocolly/colly"
)

type record struct {
	wins   int64
	losses int64
}

type mlb map[string]record

func getMLBScraped() mlb {
	t := mlb{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"),
	)
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		e.ForEach("td", func(_ int, el *colly.HTMLElement) {
			attr := el.Attr("data-label")
			switch attr {
			case "Team":
				var wins, _ = strconv.ParseInt(e.ChildText(`[data-label="Wins"]`), 10, 32)
				var losses, _ = strconv.ParseInt(e.ChildText(`[data-label="Losses"]`), 10, 32)
				r := record{
					wins:   wins,
					losses: losses,
				}
				t[el.Text] = r
			}
		})
	})

	c.Visit("https://erikberg.com/mlb/standings")
	return t
}

type mlbStandings struct {
	StandingsDate string `json:"standings_date"`
	Standing      []struct {
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
	} `json:"standing"`
}

func getMLBAPI(token string) mlb {
	client := &http.Client{}
	mlbURL, _ := url.Parse("https://erikberg.com/mlb/standings.json")
	accessHeader := fmt.Sprintf("Bearer %v", token)
	req := &http.Request{
		Method: http.MethodGet,
		URL:    mlbURL,
		Header: http.Header{
			"Authorization": []string{accessHeader},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	var rawStandings mlbStandings
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&rawStandings)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	out := mlb{}
	for _, team := range rawStandings.Standing {
		out[team.FirstName] = record{
			wins:   team.Won,
			losses: team.Lost,
		}
	}

	return out
}
