package player

import (
	"context"
	"github.com/ejagombar/spot/authstore"
	"github.com/ejagombar/spot/prechecks"
	"github.com/spf13/cobra"
)

// skip  NextCmd represents the next command
var NextCmd = &cobra.Command{
	Use:        "next",
	Short:      "Play the next song",
	Long:       ``,
	Aliases:    []string{"n"},
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
