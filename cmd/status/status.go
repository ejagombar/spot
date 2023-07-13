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
	completedHead   string
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

		songNameLine := "Song: " + fulltrack.Name
		albumNameLine := "Album: " + fulltrack.Album.Name

		var artistNameLine string
		artistList := fulltrack.Artists[0].Name

		for i := 1; i < len(fulltrack.Artists); i++ {
			artistList += ", " + fulltrack.Artists[i].Name
		}

		if len(fulltrack.Artists) > 1 {
			artistNameLine = "Artists: " + artistList
		} else {
			artistNameLine = "Artist: " + artistList
		}

		fmt.Println(songNameLine)
		fmt.Println(albumNameLine)
		fmt.Println(artistNameLine)

		minLength := 35

		barLength := findMinInt([]int{len(songNameLine), len(albumNameLine), len(artistNameLine), minLength})

		progressStamp := secondsToTimestamp(status.Progress / 1000)
		totalStamp := secondsToTimestamp(fulltrack.Duration / 1000)

		style := styleConfig{
			startChar:       " |",
			endChar:         "| ",
			completedChar:   '-',
			completedHead:   "",
			uncompletedChar: ' ',
			length:          barLength,
			frontText:       progressStamp,
			backText:        totalStamp}
		StaticProgressBar(style, status.Progress, fulltrack.Duration)
	},
}

func findMinInt(values []int) (min int) {
	min = values[0]
	for _, v := range values {
		if v > min {
			min = v
		}
	}
	return min
}

func secondsToTimestamp(secondsIn int) string {
	minutes := secondsIn / 60
	seconds := secondsIn % 60
	str := fmt.Sprintf("%d:%02d", minutes, seconds)
	return str
}

func StaticProgressBar(style styleConfig, progress int, total int) {
	paddingLength := len(style.endChar + style.startChar + style.frontText + style.backText)
	barLength := style.length - paddingLength
	percentage := int((float32(progress)/float32(total))*float32(barLength) + 0.5)
	fmt.Print(style.frontText + style.startChar)

	remainingLength := barLength - percentage

	if style.completedHead == "" {
		for i := 0; i < percentage; i++ {
			fmt.Print(string(style.completedChar))
		}
	} else {
		if percentage-len(style.completedHead) <= 0 {
			remainingLength -= len(style.completedHead)
		}
		for i := 0; i < percentage-len(style.completedHead); i++ {
			fmt.Print(string(style.completedChar))
		}
		fmt.Print(string(style.completedHead))
	}

	for i := 0; i < remainingLength; i++ {
		fmt.Print(string(style.uncompletedChar))
	}

	fmt.Print(style.endChar + style.backText + "\n")
}

func init() {
}
