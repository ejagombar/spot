package player

import (
	"context"
	"fmt"
	"github.com/ejagombar/CLSpotify/authStore"
	"github.com/spf13/cobra"
)

// skipCmd represents the skip command
var NextCmd = &cobra.Command{
	Use:   "next",
	Short: "Play the next song",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authStore.GetClient()
		if err != nil {
			fmt.Println("Unable to skip. Error:", err)
		}
		client.Next(context.Background())
	},
}

func init() {
}
