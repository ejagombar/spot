/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package player

import (
	"fmt"

	"github.com/spf13/cobra"
)

// backCmd represents the back command
var BackCmd = &cobra.Command{
	Use:   "back",
	Short: "Skip back to the previous song",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("back called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
