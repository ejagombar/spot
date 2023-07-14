package play

import (
	"context"
	"errors"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// pauseCmd represents the pause command
var AlbumCmd = &cobra.Command{
	Use:   "album",
	Short: "Specifies the search for albums",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)

		deviceID, err := SelectDevice(client)
		cobra.CheckErr(err)

		if len(args) > 0 {
			searchAlbumAndPlay(client, deviceID, ConcatArgs(args))
		} else {
			err = errors.New("Song name not provided")
		}
		cobra.CheckErr(err)
	},
}

func searchAlbumAndPlay(client *spotify.Client, deviceid spotify.ID, albumName string) (err error) {
	result, err := client.Search(context.Background(), albumName, spotify.SearchType(spotify.SearchTypeAlbum), spotify.Limit(1))
	if err != nil {
		return fmt.Errorf("Error while searching for album: %w", err)
	}

	uri := result.Albums.Albums[0].URI
	opts := spotify.PlayOptions{DeviceID: &deviceid, PlaybackContext: &uri}
	err = client.PlayOpt(context.Background(), &opts)

	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	}
	return nil
}

func init() {
}
