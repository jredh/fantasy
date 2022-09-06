package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	season = 2022
)

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

type SleeperStats struct {
	Rush_yd              float32 `json:"rush_yd"`              //: 32.0,
	Rush_fd              float32 `json:"rush_fd"`              //: 3.2,
	Rush_att             float32 `json:"rush_att"`             //: 5.0,
	Rec_yd               float32 `json:"rec_yd"`               //: 1469.0,
	Rec_td               float32 `json:"rec_td"`               //: 10.0,
	Rec_fd               float32 `json:"rec_fd"`               //: 146.9,
	Rec_5_9              float32 `json:"rec_5_9"`              //: 22.2,
	Rec_40p              float32 `json:"rec_40p"`              //: 11.1,
	Rec_30_39            float32 `json:"rec_30_39"`            //: 11.1,
	Rec_2pt              float32 `json:"rec_2pt"`              //: 1.0,
	Rec_20_29            float32 `json:"rec_20_29"`            //: 22.2,
	Rec_10_19            float32 `json:"rec_10_19"`            //: 33.3,
	Rec_0_4              float32 `json:"rec_0_4"`              //: 22.2,
	Rec                  float32 `json:"rec"`                  //: 111.0,
	Pts_std              float32 `json:"pts_std"`              //: 210.1,
	Pts_ppr              float32 `json:"pts_ppr"`              //: 321.1,
	Pts_half_ppr         float32 `json:"pts_half_ppr"`         //: 265.6,
	Gp                   float32 `json:"gp"`                   //: 18.0,
	Fum_lost             float32 `json:"fum_lost"`             //: 1.0,
	Bonus_rec_wr         float32 `json:"bonus_rec_wr"`         //: 111.0,
	Adp_std              float32 `json:"adp_std"`              //: 7.6,
	Adp_rookie           float32 `json:"adp_rookie"`           //: 999.0,
	Adp_ppr              float32 `json:"adp_ppr"`              //: 5.4,
	Adp_idp              float32 `json:"adp_idp"`              //: 7.6,
	Adp_half_ppr         float32 `json:"adp_half_ppr"`         //: 6.3,
	Adp_dynasty_std      float32 `json:"adp_dynasty_std"`      //: 6.0,
	Adp_dynasty_ppr      float32 `json:"adp_dynasty_ppr"`      //: 5.9,
	Adp_dynasty_half_ppr float32 `json:"adp_dynasty_half_ppr"` //: 6.6,
	Adp_dynasty_2qb      float32 `json:"adp_dynasty_2qb"`      //: 11.1,
	Adp_dynasty          float32 `json:"adp_dynasty"`          //: 6.3,
	Adp_2qb              float32 `json:"adp_2qb"`              //: 7.0
}

type SleeperPlayer struct {
	Years_exp         float32   `json:"years_exp"`         //: 5,
	Team              string    `json:"team"`              //: "LAR",
	Position          string    `json:"position"`          //: "WR",
	News_updated      float32   `json:"news_updated"`      //: 1661630457442,
	Metadata          *struct{} `json:"metadata"`          //: null,
	Last_name         string    `json:"last_name"`         //: "Kupp",
	Injury_status     *string   `json:"injury_status"`     //: null,
	Injury_start_date *string   `json:"injury_start_date"` //: null,
	Injury_notes      *string   `json:"injury_notes"`      //: null,
	Injury_body_part  *string   `json:"injury_body_part"`  //: null,
	First_name        string    `json:"first_name"`        //: "Cooper",
	Positions         []string  `json:"fantasy_positions"`
}
type SleeperRow struct {
	Week        *string        `json:"week"`
	Team        *string        `json:"team"`
	Stats       *SleeperStats  `json:"stats"`
	Sport       *string        `json:"sport"`       //: "nfl",
	Season_type *string        `json:"season_type"` //: "regular",
	Season      *string        `json:"season"`      //: "2022",
	Player_id   *string        `json:"player_id"`   //: "4039",
	Player      *SleeperPlayer `json:"player"`      //:,
	Opponent    *string        `json:"opponent"`    //: null,
	Game_id     *string        `json:"game_id"`     //: "season",
	Date        *string        `json:"date"`        //: null,
	Company     *string        `json:"company"`     //: "rotowire",
	Category    *string        `json:"category"`    //: "proj"
}

func main() {
	ff, err := os.Open(fmt.Sprintf(
		"./%v/%v/%v.json", season, "ffballers", "udk-data",
	))
	if err != nil {
		log.Fatal(err)
	}
	defer ff.Close()

	b, _ := ioutil.ReadAll(ff)
	var ballersData FantasyFootballersData
	err = json.Unmarshal(b, &ballersData)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range ballersData.Projections[:10] {
		fmt.Println(row)
	}

	ss, err := os.Open(fmt.Sprintf(
		"./%v/%v/%v.json", season, "sleeper", "sleeper",
	))
	if err != nil {
		log.Fatal(err)
	}
	defer ss.Close()

	b, _ = ioutil.ReadAll(ss)
	var sleeperData []SleeperRow
	err = json.Unmarshal(b, &sleeperData)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range sleeperData[:5] {
		fmt.Println(row)
	}

}
