package findmusic

import (
	"context"
	"errors"
	"fmt"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/common"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// pauseCmd represents the pause command
var PlaylistCmd = &cobra.Command{
	Use:     "playlist",
	Short:   "Specifies the search for albums",
	Long:    ``,
	Aliases: []string{"plist"},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)

		deviceID, err := common.SelectDevice(client)
		cobra.CheckErr(err)

		if len(args) > 0 {
			searchPlaylistAndPlay(client, deviceID, common.ConcatArgs(args))
		} else {
			err = errors.New("Playlist name not provided")
		}
		cobra.CheckErr(err)
	},
}

func searchPlaylistAndPlay(client *spotify.Client, deviceID spotify.ID, playlistName string) (err error) {
	err = common.SaveUserPlaylists(client, deviceID, playlistName)
	if err != nil {
		return fmt.Errorf("Error while retreiving playlist: %w", err)
	}

	bestmatch, err := common.SearchForPlaylist(playlistName)
	if err != nil {
		return fmt.Errorf("Error while searching for playlist: %w", err)
	}
	opts := spotify.PlayOptions{DeviceID: &deviceID, PlaybackContext: &bestmatch.URI}
	err = client.PlayOpt(context.Background(), &opts)

	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	}
	return nil
}

func init() {
}
