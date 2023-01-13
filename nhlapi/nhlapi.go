package nhlapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const NHLAPI = "https://statsapi.web.nhl.com/api/v1"

// type Schedule struct {
// }

// func schedule(teamID int) {
// 	today()
// }

func Today() ScheduleAPI {
	var result ScheduleAPI
	scheduleJSON := standardRequest("/schedule")
	json.Unmarshal(scheduleJSON, &result)
	return result
}

func standardRequest(endpoint string) []byte {
	resp, err := http.Get(NHLAPI + endpoint)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if json.Valid(responseData) {
		return responseData
	}
	return nil
}
