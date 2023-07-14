package play

import (
	"context"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

// pauseCmd represents the pause command
var SongCmd = &cobra.Command{
	Use:   "song",
	Short: "Specifies the search for songs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		client.Pause(context.Background())
	},
}

func init() {
}
