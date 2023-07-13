package player

import (
	"context"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// playCmd represents the play command
var PlayCmd = &cobra.Command{
	Use:   "play",
	Short: "Plays the current song",
	Long: `Plays the current song that is on the player.
        If the song is already playing or there is no song in the player, then nothing will happen.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authStore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		playerDevice, _ := client.PlayerDevices(context.Background())
		var test spotify.ID
		for _, x := range playerDevice {
			fmt.Println("Device: ")
			fmt.Println(x.ID)
			fmt.Println(x.Active)
		}
		// test = "4c3d363bf10d7394fdb1ed924359604508387da8"
		test = "71910c9ed689f71f4e8724883615d4661d717700"
		opts := spotify.PlayOptions{DeviceID: &test}
		client.PlayOpt(context.Background(), &opts)
	},
}

func init() {
}
