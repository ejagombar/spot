package common

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
)

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
