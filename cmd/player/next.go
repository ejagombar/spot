package player

import (
	"context"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

// skipCmd represents the skip command
var NextCmd = &cobra.Command{
	Use:        "next",
	Short:      "Play the next song",
	Long:       ``,
	SuggestFor: []string{"skip", "forward"},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authStore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		client.Next(context.Background())
	},
}

func init() {
}
