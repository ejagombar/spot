package cmd

import (
	"os"

	"github.com/ejagombar/spot/cmd/add"
	"github.com/ejagombar/spot/cmd/config"
	"github.com/ejagombar/spot/cmd/findmusic"
	"github.com/ejagombar/spot/cmd/info"
	"github.com/ejagombar/spot/cmd/login"
	"github.com/ejagombar/spot/cmd/player"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
)

// rootCmd represents the base command when called without any subcommands
var (
	client  *spotify.Client
	rootCmd = &cobra.Command{
		Use:   "spot",
		Short: "A brief description of your application",
		Long: `spot is a CLI tool to control your spotify account and playback.

To get started run 'spot login --help' and read the instructions to login with your account.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Add all the subcommands to the root command
func addSubCommands() {
	rootCmd.AddCommand(player.PlayCmd)
	rootCmd.AddCommand(player.PauseCmd)
	rootCmd.AddCommand(player.NextCmd)
	rootCmd.AddCommand(player.BackCmd)
	rootCmd.AddCommand(player.ShuffleCmd)

	rootCmd.AddCommand(findmusic.SongCmd)
	rootCmd.AddCommand(findmusic.AlbumCmd)
	rootCmd.AddCommand(findmusic.ArtistCmd)
	rootCmd.AddCommand(findmusic.PlaylistCmd)

	rootCmd.AddCommand(add.AddCmd)

	rootCmd.AddCommand(config.ConfigCmd)

	rootCmd.AddCommand(login.LoginCmd)

	rootCmd.AddCommand(info.InfoCmd)
}

func init() {
	initConfig()

	viper.SetDefault("auth.client_id", "")
	viper.SetDefault("auth.client_secret", "")
	viper.SetDefault("token.access", "")
	viper.SetDefault("token.refresh", "")
	viper.SetDefault("token.timeout", "")
	viper.SetDefault("config.defaultdeviceid", "")
	viper.SetDefault("appearance.status.bar.startstring", " │")
	viper.SetDefault("appearance.status.bar.endstring", "│ ")
	viper.SetDefault("appearance.status.bar.completedchar", "▒")
	viper.SetDefault("appearance.status.bar.completedhead", "░")
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

	viper.SetConfigName(".spot.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(home)
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.WriteConfigAs(home + "/.clspot.json")
		}
		cobra.CheckErr(err)
	}
	viper.ReadInConfig()
}
