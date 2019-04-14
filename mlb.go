package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gocolly/colly"
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

func getMLBScraped() mlbStandings {
	t := mlbStandings{
		Standing: make(map[string]mlbTeam),
	}

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
				r := mlbTeam{
					Won:  wins,
					Lost: losses,
				}
				t.Standing[el.Text] = r
			}
		})
	})

	c.Visit("https://erikberg.com/mlb/standings")
	return t
}

type mlbAPIStandings struct {
	StandingsDate string    `json:"standings_date"`
	Standing      []mlbTeam `json:"standing"`
}

func getMLBAPI(token string) (mlbStandings, error) {
	out := mlbStandings{
		Standing: make(map[string]mlbTeam),
	}
	rawStandings, err := compressedCall(token)
	if err != nil {
		return out, nil
	}

	for _, team := range rawStandings.Standing {
		out.Standing[team.TeamID] = team
	}
	return out, nil
}

func compressedCall(token string) (mlbAPIStandings, error) {
	client := &http.Client{}
	mlbURL, _ := url.Parse("https://erikberg.com/mlb/standings.json")
	req := &http.Request{
		Method: http.MethodGet,
		URL:    mlbURL,
		Header: http.Header{
			"Accept-Encoding": []string{"gzip"},
			"Authorization":   []string{fmt.Sprintf("Bearer %s", token)},
			"User-Agent":      []string{"MyRobot/1.0 (email@example.com)"},
		},
		Close: true,
	}

	out := mlbAPIStandings{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return out, err
	}
	defer resp.Body.Close()

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
