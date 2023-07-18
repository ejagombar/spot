package player

import (
	"context"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

// NextCmd represents the next command
var NextCmd = &cobra.Command{
	Use:        "next",
	Short:      "Play the next song",
	Long:       ``,
	SuggestFor: []string{"skip", "forward"},
	Run:        next,
}

func next(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	prechecks.DeviceAvailable(client)
	cobra.CheckErr(err)
	client.Next(context.Background())
}

func init() {
}
