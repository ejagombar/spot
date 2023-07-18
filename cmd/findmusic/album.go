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

// AlbumCmd represents the album command
var AlbumCmd = &cobra.Command{
	Use:   "album",
	Short: "Specifies the search for albums",
	Long: `Utilises the spotify search to find and play an album.

The album command prioritises playing albums over individual songs.
If set, the album will be played on the default device if it is available.
If not, the music will play on the most recently connected device.
Configure the default device using the 'config' command`,
	Example: `spot album modern life is rubbish`,
	Run:     album,
}

func album(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	cobra.CheckErr(err)
	prechecks.DeviceAvailable(client)

	deviceID, err := common.SelectDevice(client)
	cobra.CheckErr(err)

	if len(args) > 0 {
		searchAlbumAndPlay(client, deviceID, common.ConcatArgs(args))
	} else {
		err = errors.New("Song name not provided")
	}
	cobra.CheckErr(err)
}

// Search for an album and play the top result using the provided spotify client, device ID and album name.
// If unsuccessful, the errors will be wrapped and returned to the caller.
func searchAlbumAndPlay(client *spotify.Client, deviceid spotify.ID, albumName string) (err error) {
	result, err := client.Search(context.Background(), albumName, spotify.SearchType(spotify.SearchTypeAlbum), spotify.Limit(1))
	if err != nil {
		return fmt.Errorf("Error while searching for album: %w", err)
	}

	album := result.Albums.Albums[0]
	opts := spotify.PlayOptions{DeviceID: &deviceid, PlaybackContext: &album.URI}
	err = client.PlayOpt(context.Background(), &opts)

	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	} else {
		fmt.Println("Playing album " + album.Name + " by " + album.Artists[0].Name)
	}
	return nil
}

func init() {
}
