package cmd

import (
	// "fmt"
	"os"

	"github.com/ejagombar/CLSpotify/cmd/add"
	"github.com/ejagombar/CLSpotify/cmd/config"
	"github.com/ejagombar/CLSpotify/cmd/findmusic"
	"github.com/ejagombar/CLSpotify/cmd/info"
	"github.com/ejagombar/CLSpotify/cmd/login"
	"github.com/ejagombar/CLSpotify/cmd/player"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "CLSpotify",
		Short: "A brief description of your application",
		Long:  `CLSpotify is a CLI tool to control your spotify account and playback`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
	client *spotify.Client
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(player.PlayCmd)
	rootCmd.AddCommand(player.PauseCmd)
	rootCmd.AddCommand(player.NextCmd)
	rootCmd.AddCommand(player.BackCmd)
	rootCmd.AddCommand(player.ShuffleCmd)

	rootCmd.AddCommand(add.AddCmd)

	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(login.LoginCmd)
	rootCmd.AddCommand(info.InfoCmd)

	rootCmd.AddCommand(findmusic.SongCmd)
	rootCmd.AddCommand(findmusic.AlbumCmd)
	rootCmd.AddCommand(findmusic.ArtistCmd)
	rootCmd.AddCommand(findmusic.PlaylistCmd)
}

func preChecks() {}

func init() {
	initConfig()

	viper.SetDefault("auth.client_id", "")
	viper.SetDefault("auth.client_secret", "")
	viper.SetDefault("token.access", "")
	viper.SetDefault("token.refresh", "")
	viper.SetDefault("token.timeout", "")
	viper.SetDefault("config.defaultdeviceid", "")
	viper.SetDefault("appearance.status.bar.startstring", " [")
	viper.SetDefault("appearance.status.bar.endstring", "] ")
	viper.SetDefault("appearance.status.bar.completedchar", "=")
	viper.SetDefault("appearance.status.bar.completedhead", "")
	viper.SetDefault("appearance.status.bar.uncompletedchar", " ")
	viper.SetDefault("appearance.status.bar.minimumlength", 35)
	viper.SetDefault("myplaylists.items", "")
	viper.SetDefault("myplaylists.length", 0)

	viper.WriteConfig()

	addSubCommands()
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.SetConfigName(".clspot.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(home)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.WriteConfigAs(home + "/.clspot.json")
		}
		cobra.CheckErr(err)
	}
	viper.ReadInConfig()
}
