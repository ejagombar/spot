package config

import (
	"context"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// albumCmd represents the album command
var (
	defaultDevice bool
	ConfigCmd     = &cobra.Command{
		Use:   "config",
		Short: "Configure defaults and settings",
		Long:  ``,
		Run:   config,
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
				viper.Set("config.defaultDeviceID", device.ID.String())
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
	ConfigCmd.Flags().BoolVar(&defaultDevice, "SetDefaultDevice", false, "Sets the spotify device currently playing as the default")
	// ConfigCmd.Flags().String("NameDevice","",")
}
