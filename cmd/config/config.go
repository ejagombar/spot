package config

import (
	"context"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigCmd represents the config command
var (
	defaultDevice bool
	ConfigCmd     = &cobra.Command{
		Use:   "config",
		Short: "Configure defaults and settings",
		Long: `Configure some defualts and settings.
        
For security reasons and ease of use, most configuration settings are only editable through the .spot.json config file.
This file can be found by default in the home directory. It can also be created in the spot installation directory.`,
		Run: config,
	}
)

func config(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	if defaultDevice {
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		playerDevices, _ := client.PlayerDevices(context.Background())
		for _, device := range playerDevices {
			if device.Active == true {
				viper.Set("config.defaultdeviceid", device.ID.String())
				err := viper.WriteConfig()
				cobra.CheckErr(err)
				fmt.Println("Success: '" + device.Name + "' set as the default device")
				break
			}
		}
		if len(playerDevices) < 1 {
			fmt.Println("Error: No devices available")
		}
	}
}

func init() {
	ConfigCmd.Flags().BoolVar(&defaultDevice, "SetDefaultDevice", false, "Sets the spotify device that is currently active as the default when playing music")
}
