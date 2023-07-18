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
var (
	queueSong bool
	SongCmd   = &cobra.Command{
		Use:   "song",
		Short: "Specifies the search for songs",
		Long: `
If set, the song will be played on the default device if it is available.
If not, the music will play on the most recently connected device.
Configure the default device using the 'config' command`,
		Example: "spot song lucky man",
		Run:     song,
	}
)

func song(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	prechecks.DeviceAvailable(client)

	deviceID, err := common.SelectDevice(client)
	cobra.CheckErr(err)

	if len(args) > 0 {
		SearchSongAndPlay(client, deviceID, common.ConcatArgs(args), queueSong)
	} else {
		err = errors.New("Song name not provided")
	}
	cobra.CheckErr(err)
}

// Search for a song and play the top result using the provided spotify client, device ID and song name.
// If unsuccessful, the errors will be wrapped and returned to the caller.
func SearchSongAndPlay(client *spotify.Client, deviceid spotify.ID, songname string, queueSong bool) (err error) {
	result, err := client.Search(context.Background(), songname, spotify.SearchType(spotify.SearchTypeTrack), spotify.Limit(1))
	if err != nil {
		return fmt.Errorf("Error while searching for song: %w", err)
	}

	track := result.Tracks.Tracks[0]
	opts := spotify.PlayOptions{DeviceID: &deviceid, URIs: []spotify.URI{track.URI}}
	if queueSong {
		err = client.QueueSongOpt(context.Background(), result.Tracks.Tracks[0].ID, &opts)
	} else {
		err = client.PlayOpt(context.Background(), &opts)
	}
	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	} else {
		fmt.Println("Playing song " + track.Name + " by " + track.Artists[0].Name)
	}
	return nil
}

func init() {
	SongCmd.Flags().BoolVarP(&queueSong, "queue", "q", false, "If this flag is set, the song will be added to the queue")
}
