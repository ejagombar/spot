package player

import (
	"context"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
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
		client.Play(context.Background())
	},
}

func init() {
}
