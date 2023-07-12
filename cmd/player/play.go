package player

import (
	"context"
	"fmt"
	"github.com/ejagombar/CLSpotify/authStore"
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
		if err != nil {
			fmt.Println("Unable to play. Error:", err)
		}
		client.Play(context.Background())
	},
}

func init() {
}
