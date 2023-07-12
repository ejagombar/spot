package player

import (
	"context"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

// backCmd represents the back command
var PrevCmd = &cobra.Command{
	Use:   "prev",
	Short: "Skip back to the previous song",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authStore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		client.Previous(context.Background())
	},
}

func init() {
}
