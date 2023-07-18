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
var ArtistCmd = &cobra.Command{
	Use:   "artist",
	Short: "Specifies the search for albums",
	Long: `Utilises the spotify search to find and play an artist.

If set, the artist will be played on the default device if it is available.
If not, the music will play on the most recently connected device.
Configure the default device using the 'config' command`,

	Example: `spot artist the war on drugs`,
	Run:     artist,
}

func artist(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	prechecks.DeviceAvailable(client)

	deviceID, err := common.SelectDevice(client)
	cobra.CheckErr(err)

	if len(args) > 0 {
		searchAlbumAndPlay(client, deviceID, common.ConcatArgs(args))
	} else {
		err = errors.New("Artist name not provided")
	}
	cobra.CheckErr(err)
}

// Search for an artist and play the top result using the provided spotify client, device ID and artist name.
// If unsuccessful, the errors will be wrapped and returned to the caller.
func searchArtistAndPlay(client *spotify.Client, deviceid spotify.ID, artistName string) (err error) {
	result, err := client.Search(context.Background(), artistName, spotify.SearchType(spotify.SearchTypeArtist), spotify.Limit(1))
	if err != nil {
		return fmt.Errorf("Error while searching for artist: %w", err)
	}

	artist := result.Artists.Artists[0]
	opts := spotify.PlayOptions{DeviceID: &deviceid, PlaybackContext: &artist.URI}
	err = client.PlayOpt(context.Background(), &opts)

	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	} else {
		fmt.Println("Playing artist " + artist.Name)
	}
	return nil
}

func init() {
}
