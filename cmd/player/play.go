package player

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// playCmd represents the play command
var (
	Client  *spotify.Client
	num     int
	PlayCmd = &cobra.Command{
		Use:   "play",
		Short: "Plays the current song",
		Long: `Plays the current song that is on the player.
        If the song is already playing or there is no song in the player, then nothing will happen.`,
		Run: func(cmd *cobra.Command, args []string) {
			// play()
			PassClient(nil)
		},
	}
)

func play() {
	Client.Play(context.TODO())
}
func GiveNum() {
	num = 5
}

func PassClient(clientIn *spotify.Client) {
	if clientIn != nil {
		Client = clientIn
		fmt.Print("copied")
	}
	if Client != nil {
		userIn, err := Client.CurrentUser(context.TODO())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(userIn.Followers.Count)
	} else {
		fmt.Println("client is nil")
	}
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
