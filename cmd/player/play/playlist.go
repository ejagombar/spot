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
var PlaylistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "Specifies the search for albums",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)

		deviceID, err := SelectDevice(client)
		cobra.CheckErr(err)

		if len(args) > 0 {
			getUserPlaylists(client, deviceID, ConcatArgs(args))
		} else {
			err = errors.New("Playlist name not provided")
		}
		cobra.CheckErr(err)
	},
}

type playlist struct {
	Name string
	URI  string
}

func getUserPlaylists(client *spotify.Client, deviceid spotify.ID, playlistName string) (error error) {
	privateUser, err := client.CurrentUser(context.Background())
	if err != nil {
		return fmt.Errorf("Error while retrieving user data: %w", err)
	}

	userID := privateUser.ID
	results, err := client.GetPlaylistsForUser(context.Background(), userID, spotify.Offset(0))

	spotifyTotal := results.Total
	totalDownloaded := 0
	fmt.Println("total", spotifyTotal)
	for totalDownloaded < spotifyTotal {
		// matrixSize := results.Total
		// matrix := make([]playlist, matrixSize)
		length := results.Limit
		totalDownloaded += length
		fmt.Println("Next: ", results.Next)
		fmt.Println("Previous: ", results.Previous)
		fmt.Println("totaldownloaded: ", totalDownloaded)
		fmt.Println("length: ", results.Offset)
		for i := 0; i < len(results.Playlists); i++ {
			fmt.Println(i, results.Playlists[i].Name)
		}
		results, err = client.GetPlaylistsForUser(context.Background(), userID, spotify.Offset(totalDownloaded))

	}
	return nil
}

func searchPlaylistAndPlay(client *spotify.Client, deviceid spotify.ID, albumName string) (err error) {
	result, err := client.Search(context.Background(), albumName, spotify.SearchType(spotify.SearchTypePlaylist), spotify.Limit(1), spotify.Fields("Playlist(Playlist)"))
	if err != nil {
		return fmt.Errorf("Error while searching for playlist: %w", err)
	}

	uri := result.Playlists.Playlists[0].URI
	opts := spotify.PlayOptions{DeviceID: &deviceid, PlaybackContext: &uri}
	err = client.PlayOpt(context.Background(), &opts)

	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	}
	return nil
}

func init() {
}
