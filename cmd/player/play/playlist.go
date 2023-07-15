package play

import (
	"context"
	"errors"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			searchPlaylistAndPlay(client, deviceID, ConcatArgs(args))
		} else {
			err = errors.New("Playlist name not provided")
		}
		cobra.CheckErr(err)
	},
}

type playlist struct {
	Name string      `json:"name"`
	URI  spotify.URI `json:"uri"`
}

func extractStrings(slice []playlist) []string {
	result := make([]string, 0, len(slice))

	for _, item := range slice {
		result = append(result, item.Name)
	}

	return result
}

func SaveUserPlaylists(client *spotify.Client, deviceid spotify.ID, playlistName string) (error error) {
	privateUser, err := client.CurrentUser(context.Background())
	if err != nil {
		return fmt.Errorf("Error while retrieving user data: %w", err)
	}

	userID := privateUser.ID
	results, err := client.GetPlaylistsForUser(context.Background(), userID, spotify.Offset(0))
	if err != nil {
		return fmt.Errorf("Error while retrieving playlists: %w", err)
	}

	// Bug that I cba to fix: If a playlist name is changed, this will not pick up on it, as it only checks number of playlists
	if viper.GetInt("myplaylists.length") == results.Total {
		return nil
	}

	totalDownloaded := 0
	matrix := make([]playlist, results.Total)

	for do := true; do; do = (results.Next != "") {
		results, err = client.GetPlaylistsForUser(context.Background(), userID, spotify.Offset(totalDownloaded))
		if err != nil {
			return fmt.Errorf("Error while retrieving playlists: %w", err)
		}

		length := len(results.Playlists)

		for i := 0; i < length; i++ {
			matrix[i+totalDownloaded].Name = results.Playlists[i].Name
			matrix[i+totalDownloaded].URI = results.Playlists[i].URI

		}
		totalDownloaded += length
	}

	viper.Set("myplaylists.items", matrix)
	viper.Set("myplaylists.length", totalDownloaded)
	viper.WriteConfig()
	return nil
}

func loadUserPlaylists() (playlists []playlist, err error) {
	storedLength := viper.GetInt("myplaylists.length")
	playlists = make([]playlist, storedLength)
	err = viper.UnmarshalKey("myplaylists.items", &playlists)
	if err != nil {
		return nil, fmt.Errorf("Error loading user playlist data %w", err)
	}
	return playlists, nil
}

func searchPlaylistAndPlay(client *spotify.Client, deviceID spotify.ID, playlistName string) (err error) {
	err = SaveUserPlaylists(client, deviceID, playlistName)
	if err != nil {
		return fmt.Errorf("Error while searching for playlist: %w", err)
	}

	playlists, _ := loadUserPlaylists()
	bagsize := []int{2}
	cm := closestmatch.New(extractStrings(playlists), bagsize)

	fmt.Println(cm.Closest(playlistName))
	if err != nil {
		return fmt.Errorf("Error while attempting to play: %w", err)
	}
	return nil
}

func init() {
}
