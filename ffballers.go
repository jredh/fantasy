package main

type FantasyFootballerRow struct {
	PlayerId                      string                                                   `json:"player_id"`  // "player_id": "732",
	AnalystId                     string                                                   `json:"analyst_id"` // "analyst_id": "1",
	Season                        string                                                   `json:"season"`     // "season": "2022",
	Week                          string                                                   `json:"week"`       // "week": "1",
	Passing_attempts              string/*float32*/ `json:"passing_attempts"`              // "passing_attempts": "0.00",
	Passing_yards                 string/*float32*/ `json:"passing_yards"`                 // "passing_yards": "269.39",
	Passing_completions           string/*float32*/ `json:"passing_completions"`           // "passing_completions": "0.00",
	Passing_touchdowns            string/*float32*/ `json:"passing_touchdowns"`            // "passing_touchdowns": "1.56",
	Rushing_attempts              string/*float32*/ `json:"rushing_attempts"`              // "rushing_attempts": "0.00",
	Rushing_yards                 string/*float32*/ `json:"rushing_yards"`                 // "rushing_yards": "5.37",
	Rushing_yards_per_attempt     string/*float32*/ `json:"rushing_yards_per_attempt"`     // "rushing_yards_per_attempt": "0.00",
	Rushing_touchdowns            string/*float32*/ `json:"rushing_touchdowns"`            // "rushing_touchdowns": "0.09",
	Receptions                    string/*float32*/ `json:"receptions"`                    // "receptions": "0.00",
	Receiving_yards               string/*float32*/ `json:"receiving_yards"`               // "receiving_yards": "0.00",
	Receiving_yards_per_reception string/*float32*/ `json:"receiving_yards_per_reception"` // "receiving_yards_per_reception": "0.00",
	Receiving_touchdowns          string/*float32*/ `json:"receiving_touchdowns"`          // "receiving_touchdowns": "0.00",
	Interceptions_thrown          string/*float32*/ `json:"interceptions_thrown"`          // "interceptions_thrown": "0.87",
	Fumbles_lost                  string/*float32*/ `json:"fumbles_lost"`                  // "fumbles_lost": "0.28",
	Risk                          string/*float32*/ `json:"risk"`                          // "risk": "0.00",
	Created_at                    string                                                   `json:"created_at"`       // "created_at": "2022-08-31 10:30:04",
	Updated_at                    string                                                   `json:"updated_at"`       // "updated_at": "2022-09-06 14:00:09",
	Analyst_name                  string                                                   `json:"analyst_name"`     // "analyst_name": "Andy",
	Name                          string                                                   `json:"name"`             // "name": "Matt Ryan",
	Fantasy_position              string                                                   `json:"fantasy_position"` // "fantasy_position": "QB",
	Team                          string                                                   `json:"team"`             // "team": "IND",
	Number                        string/*int*/ `json:"number"`                            // "number": "2",
	Bye_week                      string/*int*/ `json:"bye_week"`                          // "bye_week": "14",
	Udk                           string/*int*/ `json:"udk"`                               // "udk": "1",
	Slug                          string                                                   `json:"slug"`              // "slug": "matt-ryan",
	Injury_status                 *string                                                  `json:"injury_status"`     // "injury_status": null,
	Injury_start_date             *string                                                  `json:"injury_start_date"` // "injury_start_date": null,
	Injury_body_part              *string                                                  `json:"injury_body_part"`  // "injury_body_part": null,
	Injury_notes                  *string                                                  `json:"injury_notes"`      // "injury_notes": null,
	Adp                           string/*float32*/ `json:"adp"`                           // "adp": "183.10",
	Adp_ppr                       string/*float32*/ `json:"adp_ppr"`                       // "adp_ppr": "185.60",
	Headshot                      string                                                   `json:"headshot"` // "headshot": "headshots\/fdbg732.jpg",
	Tagged                        bool                                                     `json:"tagged"`   // "tagged": false
}

type FantasyFootballersData struct {
	Projections []FantasyFootballerRow `json:"projections"`
}
