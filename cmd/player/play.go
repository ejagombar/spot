/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package player

import (
	"fmt"

	"github.com/spf13/cobra"
)

// playCmd represents the play command
var PlayCmd = &cobra.Command{
	Use:   "play",
	Short: "Plays the current song",
	Long: `Plays the current song that is on the player.
    If the song is already playing or there is no song in the player, then nothing will happen.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("play called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
