package main

import (
	"time"
)

type ScheduleAPI struct {
	Copyright    string `json:"copyright"`
	TotalItems   int    `json:"totalItems"`
	TotalEvents  int    `json:"totalEvents"`
	TotalGames   int    `json:"totalGames"`
	TotalMatches int    `json:"totalMatches"`
	MetaData     struct {
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
	Wait  int `json:"wait"`
	Dates []struct {
		Date         string `json:"date"`
		TotalItems   int    `json:"totalItems"`
		TotalEvents  int    `json:"totalEvents"`
		TotalGames   int    `json:"totalGames"`
		TotalMatches int    `json:"totalMatches"`
		Games        []struct {
			GamePk   int       `json:"gamePk"`
			Link     string    `json:"link"`
			GameType string    `json:"gameType"`
			Season   string    `json:"season"`
			GameDate time.Time `json:"gameDate"`
			Status   struct {
				AbstractGameState string `json:"abstractGameState"`
				CodedGameState    string `json:"codedGameState"`
				DetailedState     string `json:"detailedState"`
				StatusCode        string `json:"statusCode"`
				StartTimeTBD      bool   `json:"startTimeTBD"`
			} `json:"status"`
			Teams struct {
				Away struct {
					LeagueRecord LeagueRecord `json:"leagueRecord"`
					Score        int          `json:"score"`
					Team         Team         `json:"team"`
				} `json:"away"`
				Home struct {
					LeagueRecord LeagueRecord `json:"leagueRecord"`
					Score        int          `json:"score"`
					Team         Team         `json:"team"`
				} `json:"home"`
			} `json:"teams"`
			Venue   Venue `json:"venue"`
			Content struct {
				Link string `json:"link"`
			} `json:"content"`
		} `json:"games"`
		Events  []interface{} `json:"events"`
		Matches []interface{} `json:"matches"`
	} `json:"dates"`
}

type FranchisesAPI struct {
	Copyright  string `json:"copyright"`
	Franchises []struct {
		FranchiseID      int    `json:"franchiseId"`
		FirstSeasonID    int    `json:"firstSeasonId"`
		MostRecentTeamID int    `json:"mostRecentTeamId"`
		TeamName         string `json:"teamName"`
		LocationName     string `json:"locationName"`
		Link             string `json:"link"`
		LastSeasonID     int    `json:"lastSeasonId,omitempty"`
	} `json:"franchises"`
}

type GameType struct {
	id          string `json:"id"`
	description string `json:"description"`
	postseason  bool   `json:"postseason"`
}

type LeagueRecord struct {
	Wins  int    `json:"wins"`
	Loses int    `json:"loses"`
	Ties  int    `json:"ties"`
	Type  string `json:"type"`
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type TimeZone struct {
	ID     string `json:"id"`
	Offset int    `json:"offset"`
	Tz     string `json:"tz"`
}

type Venue struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Link     string   `json:"link"`
	City     string   `json:"city"`
	TimeZone TimeZone `json:"time_zone"`
}
