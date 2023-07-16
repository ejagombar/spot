package player

import (
	"context"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

// skipCmd represents the skip command
var ShuffleCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "Switch between shuffle and non-shuffle modes",
	Long: `This command will shuffle the queued songs by default.
    The shuffle can be disabled by adding "false or "f" after the shuffle argument`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		value := true
		if len(args) > 0 {
			if args[0] == "false" || args[0] == "f" {
				value = false
			}
		}
		err = client.Shuffle(context.Background(), value)
		cobra.CheckErr(nil)
	},
}

func init() {
}
