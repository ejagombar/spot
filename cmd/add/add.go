package add

import (
	"context"
	"fmt"

	"errors"
	"github.com/ejagombar/spot/authstore"
	"github.com/ejagombar/spot/common"
	"github.com/ejagombar/spot/playlist"
	"github.com/ejagombar/spot/prechecks"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// backCmd represents the back command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the current song to a playlist",
	Long: `Add the current song to a playlist.

Specify the playlist by typing it's name after the command.
The playlist will be fuzzy searched against the list of playlists in your library.
Spaces between words in the playlist name are accepted.`,

	Example:    `spot add my personal playlist`,
	SuggestFor: []string{"prev", "previous"},
	Run:        add,
}

func add(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	cobra.CheckErr(err)
	prechecks.DeviceAvailable(client)

	deviceID, err := common.SelectDevice(client)
	cobra.CheckErr(err)

	if len(args) > 0 {
		searchPlaylistAndAdd(client, deviceID, common.ConcatArgs(args))
	} else {
		err = errors.New("Playlist name not provided")
	}

	cobra.CheckErr(err)
}

// This functin attempts to seach for a playlist in the list of saved playlists and add the currently
// playing song to that playlist. A message will be displayed to the screen if the song is successfully
// added. The message will the user the name of the song and which playlist it has been added to. This
// is needed as the fuzzy search is not always accurate.
func searchPlaylistAndAdd(client *spotify.Client, deviceID spotify.ID, playlistName string) (err error) {
	err = playlist.SaveUserPlaylists(client, deviceID, playlistName)
	if err != nil {
		return fmt.Errorf("Error while retreiving playlist: %w", err)
	}

	bestmatch, err := playlist.SearchForPlaylist(playlistName)
	if err != nil {
		return fmt.Errorf("Error while searching for playlist: %w", err)
	}
	playerSong, err := client.PlayerCurrentlyPlaying(context.Background())
	_, err = client.AddTracksToPlaylist(context.Background(), bestmatch.ID, playerSong.Item.ID)
	cobra.CheckErr(err)

	if err != nil {
		return fmt.Errorf("Error while attempting to add to playlist: %w", err)
	} else {
		fmt.Println("Added " + playerSong.Item.Name + " to " + bestmatch.Name)
	}

	return nil
}

func init() {
}
