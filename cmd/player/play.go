package player

import (
	"context"
	"errors"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
)

// playCmd represents the play command
var PlayCmd = &cobra.Command{
	Use:   "play",
	Short: "Plays the current song",
	Long: `Plays the current song that is on the player.
        If the song is already playing or there is no song in the player, then nothing will happen.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)

		deviceID, err := selectDevice(client)
		cobra.CheckErr(err)

		opts := spotify.PlayOptions{DeviceID: &deviceID}
		client.PlayOpt(context.Background(), &opts)
	},
}

func selectDevice(client *spotify.Client) (deviceID spotify.ID, err error) {
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

func init() {
}
