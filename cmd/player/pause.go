/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package player

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pauseCmd represents the pause command
var PauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses the current song",
	Long: `Use this command to pause the current song playing
    If no song is playing, the command will not do anything.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pause called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pauseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pauseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
