package player

import (
	"context"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

// skip  NextCmd represents the next command
var NextCmd = &cobra.Command{
	Use:        "next",
	Short:      "Play the next song",
	Long:       ``,
	SuggestFor: []string{"skip", "forward"},
	Run:        next,
}

func next(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	cobra.CheckErr(err)
	prechecks.DeviceAvailable(client)
	client.Next(context.Background())
}

func init() {
}
