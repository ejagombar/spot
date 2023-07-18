package player

import (
	"context"

	"github.com/ejagombar/spot/authstore"
	"github.com/ejagombar/spot/prechecks"
	"github.com/spf13/cobra"
)

// ShuffleCmd represents the shuffle command
var ShuffleCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "Switch between shuffle and non-shuffle modes",
	Long: `This command will shuffle the queued songs by default.
    The shuffle can be disabled by adding "false or "f" after the shuffle argument`,
	Aliases: []string{"s"},
	Run:     shuffle,
}

func shuffle(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	cobra.CheckErr(err)
	prechecks.DeviceAvailable(client)
	value := true
	if len(args) > 0 {
		if args[0] == "false" || args[0] == "f" {
			value = false
		}
	}
	err = client.Shuffle(context.Background(), value)
	cobra.CheckErr(nil)
}

func init() {
}
