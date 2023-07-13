/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package status

import (
	"context"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// StatusCmd represents the album command
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display general info",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		status, err := client.PlayerCurrentlyPlaying(context.Background())
		fulltrack := status.Item

		fmt.Println("Song:", fulltrack.Name)
		fmt.Println("Album:", fulltrack.Album.Name)
		artistList := fulltrack.Artists[0].Name

		for i := 1; i < len(fulltrack.Artists); i++ {
			artistList += ", " + fulltrack.Artists[i].Name
		}
		if len(fulltrack.Artists) > 1 {
			fmt.Println("Artists:", artistList)
		} else {
			fmt.Println("Artist:", artistList)
		}

		songProgress := status.Progress
		theme := progressbar.Theme{Saucer: "-", SaucerHead: ">", AltSaucerHead: "<", SaucerPadding: " ", BarStart: "|", BarEnd: "|"}

		bar := progressbar.NewOptions(fulltrack.Duration, progressbar.OptionSetTheme(theme), progressbar.OptionSetWidth(20))
		bar.Set(songProgress)

	},
}

func init() {
}
