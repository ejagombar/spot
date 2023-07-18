package findmusic

import (
	"context"
	"errors"
	"fmt"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/common"
	playlistLib "github.com/ejagombar/CLSpotify/playlist"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// pauseCmd represents the pause command
var PlaylistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "Specifies the search for albums",
	Long: `Utilises fuzzy finding to search and play a playlist on your account

If set, the album will be played on the default device if it is available.
If not, the music will play on the most recently connected device.
Configure the default device using the 'config' command

The fuzzy finding will find the playlist with the closest name.
This allows for typos and shortened spellings however it may occasionally return the wrong result.

A list of all playlists in the connected account is saved to the config file.
This list is only updated when the a new playlist is added or deleted on the connected spotify account.
If a playlist is renamed, the list will not be updated.`,
	Aliases: []string{"plist"},
	Run:     playlist,
}

func playlist(cmd *cobra.Command, args []string) {
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
}

func searchPlaylistAndPlay(client *spotify.Client, deviceID spotify.ID, playlistName string) (err error) {
	err = playlistLib.SaveUserPlaylists(client, deviceID, playlistName)
	if err != nil {
		return fmt.Errorf("Error while retreiving playlist: %w", err)
	}

	bestmatch, err := playlistLib.SearchForPlaylist(playlistName)
	if err != nil {
		return fmt.Errorf("Error while searching for playlist: %w", err)
	}
	opts := spotify.PlayOptions{DeviceID: &deviceID, PlaybackContext: &bestmatch.URI}
	err = client.PlayOpt(context.Background(), &opts)

	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	} else {
		fmt.Println("Playing from playlist " + bestmatch.Name)
	}

	return nil
}

func init() {
}
