package main

import (
	"strconv"

	"github.com/gocolly/colly"
)

type record struct {
	wins   int64
	losses int64
}

type mlb map[string]record

func getMLB() mlb {
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
<<<<<<< Updated upstream
=======

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
	accessHeader := fmt.Sprintf("Bearer %v", token)
	req := &http.Request{
		Method: http.MethodGet,
		URL:    mlbURL,
		Header: http.Header{
			"Authorization":   []string{accessHeader},
			"Accept-Encoding": []string{"gzip"},
		},
		Close: true,
	}

	out := mlbAPIStandings{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return out, err
	}

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return out, err
	}
	b, err := ioutil.ReadAll(gz)
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(b))
	err = decoder.Decode(&out)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return out, err
	}

	return out, nil
}
>>>>>>> Stashed changes
