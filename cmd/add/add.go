package add

import (
	"context"
	"fmt"

	"errors"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/common"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// backCmd represents the back command
var AddCmd = &cobra.Command{
	Use:        "add",
	Short:      "Add the current song to a playlist",
	Long:       ``,
	SuggestFor: []string{"prev", "previous"},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)

		deviceID, err := common.SelectDevice(client)
		cobra.CheckErr(err)

		if len(args) > 0 {
			searchPlaylistAndAdd(client, deviceID, common.ConcatArgs(args))
		} else {
			err = errors.New("Playlist name not provided")
		}

		cobra.CheckErr(err)
	},
}

func searchPlaylistAndAdd(client *spotify.Client, deviceID spotify.ID, playlistName string) (err error) {
	err = common.SaveUserPlaylists(client, deviceID, playlistName)
	if err != nil {
		return fmt.Errorf("Error while retreiving playlist: %w", err)
	}

	bestmatch, err := common.SearchForPlaylist(playlistName)
	if err != nil {
		return fmt.Errorf("Error while searching for playlist: %w", err)
	}
	playerSong, err := client.PlayerCurrentlyPlaying(context.Background())
	fmt.Println(playerSong.Item.Name)
	_, err = client.AddTracksToPlaylist(context.Background(), bestmatch.ID, playerSong.Item.ID)
	cobra.CheckErr(err)

	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	}
	return nil
}

func init() {
}
