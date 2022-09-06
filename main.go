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

func openU[T any](season int, provider, source string) (*T, error) {
	ff, err := os.Open(fmt.Sprintf(
		"./%v/%v/%v.json", season, provider, source,
	))
	if err != nil {
		return nil, err
	}
	defer ff.Close()

	b, err := ioutil.ReadAll(ff)
	if err != nil {
		return nil, err
	}

	var data T
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {

	ballersData, err := openU[FantasyFootballersData](2022, "ffballers", "udk-data")
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range ballersData.Projections[:10] {
		fmt.Println(row)
	}

	sd, err := openU[SleeperData](2022, "sleeper", "sleeper")
	if err != nil {
		log.Fatal(err)
	}
	sleeperData := *sd

	for _, row := range sleeperData[:5] {
		fmt.Println(row)
	}

}
