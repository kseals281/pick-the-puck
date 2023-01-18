/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/kseals281/pick-the-puck/nhlapi"
	"time"

	"github.com/spf13/cobra"
)

// yesterdayCmd represents the yesterday command
var yesterdayCmd = &cobra.Command{
	Use:   "yesterday",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		today := time.Now()
		yesterday := today.Add(-(time.Hour * 24))
		y, m, d := yesterday.Date()
		startDate := fmt.Sprintf("startDate=%d-%d-%d", y, m, d)
		endDate := fmt.Sprintf("endDate=%d-%d-%d", y, m, d)
		schedule := nhlapi.Schedule(startDate + "&" + endDate)
		readScores(schedule.Dates[0].Games)
	},
}

func readScores(games []nhlapi.Games) {
	for _, game := range games {
		fmt.Printf("%s: %d\t@\t%s: %d\n",
			game.Teams.Away.Team.Name, game.Teams.Away.Score,
			game.Teams.Home.Team.Name, game.Teams.Home.Score)
	}
}

func init() {
	rootCmd.AddCommand(yesterdayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// yesterdayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// yesterdayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
