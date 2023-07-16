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
var SongCmd = &cobra.Command{
	Use:   "song",
	Short: "Specifies the search for songs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)

		deviceID, err := common.SelectDevice(client)
		cobra.CheckErr(err)

		if len(args) > 0 {
			SearchSongAndPlay(client, deviceID, common.ConcatArgs(args))
		} else {
			err = errors.New("Song name not provided")
		}
		cobra.CheckErr(err)
	},
}

func SearchSongAndPlay(client *spotify.Client, deviceid spotify.ID, songname string) (err error) {
	result, err := client.Search(context.Background(), songname, spotify.SearchType(spotify.SearchTypeTrack), spotify.Limit(1))
	if err != nil {
		return fmt.Errorf("Error while searching for song: %w", err)
	}

	uri := result.Tracks.Tracks[0].URI
	opts := spotify.PlayOptions{DeviceID: &deviceid, URIs: []spotify.URI{uri}}
	err = client.PlayOpt(context.Background(), &opts)

	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	}
	return nil
}

func init() {
}
