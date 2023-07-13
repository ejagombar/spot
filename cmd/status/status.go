/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package status

import (
	"context"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

type styleConfig struct {
	startChar       string
	endChar         string
	completedChar   byte
	uncompletedChar byte
	frontText       string
	backText        string
	length          int
}

// StatusCmd represents the album command
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display general info",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authstore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		status, err := client.PlayerCurrentlyPlaying(context.Background())
		fulltrack := status.Item

		fmt.Println("Song:", fulltrack.Name)
		fmt.Println("Album:", fulltrack.Album.Name)
		artistList := fulltrack.Artists[0].Name

		for i := 1; i < len(fulltrack.Artists); i++ {
			artistList += ", " + fulltrack.Artists[i].Name
		}
		if len(fulltrack.Artists) > 1 {
			fmt.Println("Artists:", artistList)
		} else {
			fmt.Println("Artist:", artistList)
		}
		// songProgress := status.Progress
		style := styleConfig{startChar: "|", endChar: "|", completedChar: '-', uncompletedChar: ' ', length: 30}
		StaticProgressBar(style, status.Progress, fulltrack.Duration)
	},
}

func StaticProgressBar(style styleConfig, progress int, total int) {
	paddingLength := len(style.endChar + style.startChar)
	barLength := style.length - paddingLength
	percentage := int((float32(progress)/float32(total))*float32(barLength) + 0.5)
	fmt.Print(style.frontText + style.startChar)
	for i := 0; i < percentage; i++ {
		fmt.Print(string(style.completedChar))
	}
	for i := 0; i < style.length-percentage; i++ {
		fmt.Print(string(style.uncompletedChar))
	}
	fmt.Print(style.endChar + style.backText + "\n")
}

func init() {
}
