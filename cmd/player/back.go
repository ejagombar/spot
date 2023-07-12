package player

import (
	"context"
	"fmt"
	"github.com/ejagombar/CLSpotify/authStore"
	"github.com/spf13/cobra"
)

// backCmd represents the back command
var PrevCmd = &cobra.Command{
	Use:   "prev",
	Short: "Skip back to the previous song",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authStore.GetClient()
		if err != nil {
			fmt.Println("Unable to skip. Error:", err)
		}
		client.Previous(context.Background())
	},
}

func init() {
}
