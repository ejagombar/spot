/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package song

import (
	"context"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

// songCmd represents the song command
var SongCmd = &cobra.Command{
	Use:   "song",
	Short: "Information about the current song",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		fmt.Println("test")
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		playerDevice, _ := client.PlayerDevices(context.Background())
		fmt.Println("is device restricted", playerDevice[0].ID)
		fmt.Println("is device restricted", playerDevice[0].Restricted)
		fmt.Println("is device active", playerDevice[0].Active)

	},
}

func init() {
}
