package prechecks

import (
	"context"
	"errors"
	"fmt"
	"github.com/zmb3/spotify/v2"
)

// Ensures that a devices logged into the user's account running spotify.
// If none are, an error is printed to the screen.
func DeviceAvailable(client *spotify.Client) {
	errMsg := "Unable to run command. Error:"

	playerDevices, err := client.PlayerDevices(context.Background())
	if err != nil {
		fmt.Println(errMsg, err)
	}
	if len(playerDevices) == 0 {
		err = errors.New("No devices available")
		fmt.Println(errMsg, err)
	}
}
