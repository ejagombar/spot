/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package player

import (
	"fmt"

	"github.com/spf13/cobra"
)

// skipCmd represents the skip command
var SkipCmd = &cobra.Command{
	Use:   "skip",
	Short: "Play the next song",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("skip called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// skipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// skipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
