/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	"os"

	// "github.com/ejagombar/CLSpotify/cmd/album"
	// "github.com/ejagombar/CLSpotify/cmd/artist"
	"github.com/ejagombar/CLSpotify/cmd/auth"
	"github.com/ejagombar/CLSpotify/cmd/config"
	"github.com/ejagombar/CLSpotify/cmd/player"
	"github.com/ejagombar/CLSpotify/cmd/status"

	// "github.com/ejagombar/CLSpotify/cmd/playlist"
	"github.com/ejagombar/CLSpotify/cmd/song"
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

	// rootCmd.AddCommand(album.AlbumCmd)
	rootCmd.AddCommand(song.SongCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	// rootCmd.AddCommand(artist.ArtistCmd)
	// rootCmd.AddCommand(playlist.PlaylistCmd)
	rootCmd.AddCommand(auth.LoginCmd)
	rootCmd.AddCommand(status.StatusCmd)
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
