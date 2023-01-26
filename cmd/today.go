/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/kseals281/pick-the-puck/nhlapi"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A list of NHL matchups on the current day",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		today := time.Now()
		schedule := nhlapi.Schedule(nhlapi.DateEndpoint(today))
		readGames(schedule.Dates[0].Games)
	},
}

func readGames(games []nhlapi.Games) {
	fmt.Println("---------------------------------------------------------")
	for _, game := range games {
		gameData := nhlapi.Game(game.Link[7:]).GameData
		teams := gameData.Teams
		loc := time.FixedZone("UTC-5", -5*60*60)
		gameTime := gameData.DateTime.DateTime.In(loc)
		fmt.Printf("|%-22s| %7v |%22s|\n", teams.Away.Name, gameTime.Format(time.Kitchen), teams.Home.Name)
		fmt.Println("---------------------------------------------------------")
	}
}

func init() {
	rootCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
