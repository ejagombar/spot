/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package playlist

import (
	"fmt"

	"github.com/spf13/cobra"
)

// playlistCmd represents the playlist command
var PlaylistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "Operate on a playlist",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("playlist called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
