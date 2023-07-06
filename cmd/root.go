/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/ejagombar/CLSpotify/cmd/album"
	"github.com/ejagombar/CLSpotify/cmd/artist"
	"github.com/ejagombar/CLSpotify/cmd/auth"
	"github.com/ejagombar/CLSpotify/cmd/player"
	"github.com/ejagombar/CLSpotify/cmd/playlist"
	"github.com/ejagombar/CLSpotify/cmd/song"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "CLSpotify",
	Short: "A brief description of your application",
	Long:  `CLSpotify is a CLI tool to control your spotify account and playback`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.CLSpotify.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(player.PlayCmd)
	rootCmd.AddCommand(player.PauseCmd)
	rootCmd.AddCommand(player.SkipCmd)
	rootCmd.AddCommand(player.BackCmd)

	rootCmd.AddCommand(album.AlbumCmd)
	rootCmd.AddCommand(song.SongCmd)
	rootCmd.AddCommand(artist.ArtistCmd)
	rootCmd.AddCommand(player.PlayCmd)
	rootCmd.AddCommand(playlist.PlaylistCmd)
	rootCmd.AddCommand(auth.AuthCmd)

}
