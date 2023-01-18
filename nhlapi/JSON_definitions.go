package nhlapi

import (
	"time"
)

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

type GameAPI struct {
	Copyright string
	GamePk    int
	Link      string
	GameData  GameData
}

type ScheduleAPI struct {
	Copyright    string `json:"copyright"`
	TotalItems   int    `json:"totalItems"`
	TotalEvents  int    `json:"totalEvents"`
	TotalGames   int    `json:"totalGames"`
	TotalMatches int    `json:"totalMatches"`
	MetaData     struct {
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
	Wait  int     `json:"wait"`
	Dates []Dates `json:"dates"`
}

/******************************************************************************
******************************************************************************/

type Dates struct {
	Date         string        `json:"date"`
	TotalItems   int           `json:"totalItems"`
	TotalEvents  int           `json:"totalEvents"`
	TotalGames   int           `json:"totalGames"`
	TotalMatches int           `json:"totalMatches"`
	Games        []Games       `json:"games"`
	Events       []interface{} `json:"events"`
	Matches      []interface{} `json:"matches"`
}

type Games struct {
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
			LeagueRecord LeagueRecord `json:"league_record"`
			Score        int          `json:"score"`
			Team         Team         `json:"team"`
		}
		Home struct {
			LeagueRecord LeagueRecord `json:"league_record"`
			Score        int          `json:"score"`
			Team         Team         `json:"team"`
		} `json:"home"`
	} `json:"teams"`
}

type GameData struct {
	Teams struct {
		Home Team
		Away Team
	}
}

type GameType struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Postseason  bool   `json:"postseason"`
}

type LeagueRecord struct {
	Wins  int    `json:"wins"`
	Loses int    `json:"loses"`
	Ties  int    `json:"ties"`
	Type  string `json:"type"`
}

type Team struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Venue        Venue  `json:"venue"`
	Abbreviation string `json:"abbreviation"`
	TriCode      string `json:"triCode"`
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
