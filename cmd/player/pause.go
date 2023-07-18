package player

import (
	"context"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

// pauseCmd represents the pause command
var PauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses the current song",
	Long: `Use this command to pause the current song playing
    If no song is playing, the command will not do anything.`,
	Run: pause,
}

func pause(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	cobra.CheckErr(err)
	prechecks.DeviceAvailable(client)
	client.Pause(context.Background())
}

func init() {
}
