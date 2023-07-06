/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package song

import (
	"fmt"

	"github.com/spf13/cobra"
)

// songCmd represents the song command
var SongCmd = &cobra.Command{
	Use:   "song",
	Short: "Information about the current song",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("song called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// songCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// songCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
