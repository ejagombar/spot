package play

import (
	"context"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// playCmd represents the play command
var PlayCmd = &cobra.Command{
	Use:   "play",
	Short: "Plays the current song",
	Long: `Plays the current song that is on the player.
        If the song is already playing or there is no song in the player, then nothing will happen.`,
	Run: play,
}

func play(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	prechecks.DeviceAvailable(client)
	cobra.CheckErr(err)

	deviceID, err := SelectDevice(client)
	cobra.CheckErr(err)

	if len(args) > 0 {
		err = SearchSongAndPlay(client, deviceID, ConcatArgs(args))
	} else {
		opts := spotify.PlayOptions{DeviceID: &deviceID}
		client.PlayOpt(context.Background(), &opts)
	}
	cobra.CheckErr(err)
}

func init() {
	PlayCmd.AddCommand(SongCmd)
	PlayCmd.AddCommand(AlbumCmd)
	PlayCmd.AddCommand(ArtistCmd)
	PlayCmd.AddCommand(PlaylistCmd)
}
