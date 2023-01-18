package nhlapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const NHLAPI = "https://statsapi.web.nhl.com/api/v1"

const (
	NJD = 1
	NYI = 2
)

func DateEndpoint(date time.Time) string {
	y, m, d := date.Date()
	startDate := fmt.Sprintf("startDate=%d-%d-%d", y, m, d)
	endDate := fmt.Sprintf("endDate=%d-%d-%d", y, m, d)
	return fmt.Sprint(startDate + "&" + endDate)
}

func Game(endpoint string) GameAPI {
	var result GameAPI
	gameJSON := standardRequest(endpoint)
	json.Unmarshal(gameJSON, &result)
	return result
}

func Schedule(endpoint string) ScheduleAPI {
	var result ScheduleAPI
	scheduleJSON := standardRequest("/schedule?" + endpoint)
	json.Unmarshal(scheduleJSON, &result)
	return result
}

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
