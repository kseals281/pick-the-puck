package main

import (
	"bufio"
	"encoding/json"
	"net/http"
)

const NHLAPI = "https://statsapi.web.nhl.com/api/v1"

// type Schedule struct {
// }

// func schedule(teamID int) {
// 	today()
// }

// func today() {
// 	resp, _ := http.Get(NHLAPI + "schedule")
// 	fmt.Println(resp)
// }

func standardRequest(endpoint string) []byte {
	resp, _ := http.Get(NHLAPI + endpoint)
	reader := bufio.NewReader(resp.Body)
	var respJSON []byte
	reader.Read(respJSON)

	if json.Valid(respJSON) {
		return respJSON
	}
	return nil
}
