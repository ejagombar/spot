package player

import (
	"context"
	"github.com/ejagombar/spot/authstore"
	"github.com/ejagombar/spot/prechecks"
	"github.com/spf13/cobra"
)

// backCmd represents the back command
var BackCmd = &cobra.Command{
	Use:        "back",
	Short:      "Skip back to the previous song",
	Long:       ``,
	Aliases:    []string{"b"},
	SuggestFor: []string{"prev", "previous"},
	Run:        back,
}

func back(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	cobra.CheckErr(err)
	prechecks.DeviceAvailable(client)
	client.Previous(context.Background())
}

func init() {
}
