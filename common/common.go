package common

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"github.com/zmb3/spotify/v2"
	"strings"
)

type Playlist struct {
	Name string      `json:"name"`
	URI  spotify.URI `json:"uri"`
	ID   spotify.ID  `json:"ID"`
}

func SelectDevice(client *spotify.Client) (deviceID spotify.ID, err error) {
	playerDevice, _ := client.PlayerDevices(context.Background())
	defaultDevice := fmt.Sprint(viper.Get("config.defaultdeviceid"))

	if len(playerDevice) < 1 {
		return "", errors.New("No devices available")
	}

	deviceID = playerDevice[0].ID

	if defaultDevice == "" {
		return deviceID, nil
	}

	for _, device := range playerDevice {
		if device.ID.String() == defaultDevice {
			deviceID = device.ID
			break
		}
	}

	return deviceID, nil
}

func ConcatArgs(args []string) (str string) {
	str = ""
	for _, word := range args {
		str += word + " "
	}
	return str
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
	matrix := make([]Playlist, results.Total)

	for do := true; do; do = (results.Next != "") {
		results, err = client.GetPlaylistsForUser(context.Background(), userID, spotify.Offset(totalDownloaded))
		if err != nil {
			return fmt.Errorf("Error while retrieving playlists: %w", err)
		}

		length := len(results.Playlists)

		for i := 0; i < length; i++ {
			matrix[i+totalDownloaded].Name = results.Playlists[i].Name
			matrix[i+totalDownloaded].URI = results.Playlists[i].URI
			matrix[i+totalDownloaded].ID = results.Playlists[i].ID

		}
		totalDownloaded += length
	}

	viper.Set("myplaylists.items", matrix)
	viper.Set("myplaylists.length", totalDownloaded)
	viper.WriteConfig()
	return nil
}

func loadUserPlaylists() (playlists []Playlist, err error) {
	storedLength := viper.GetInt("myplaylists.length")
	playlists = make([]Playlist, storedLength)
	err = viper.UnmarshalKey("myplaylists.items", &playlists)
	if err != nil {
		return nil, fmt.Errorf("Error loading user playlist data %w", err)
	}
	return playlists, nil
}

func SearchForPlaylist(playlistName string) (bestmatch Playlist, err error) {

	playlists, _ := loadUserPlaylists()

	bestMatchIndex := -1
	bestMatchDistance := -1

	for i, item := range playlists {
		distance := levenshtein.DistanceForStrings([]rune(strings.ToLower(item.Name)), []rune(strings.ToLower(playlistName)), levenshtein.DefaultOptions)
		if bestMatchDistance == -1 || distance < bestMatchDistance {
			bestMatchDistance = distance
			bestMatchIndex = i
		}
	}

	if bestMatchIndex == -1 {
		return Playlist{}, errors.New("No match found")
	}
	return playlists[bestMatchIndex], nil
}
