/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"context"
	"errors"
	"fmt"

	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type styleConfig struct {
	startString     string
	endString       string
	completedChar   string
	completedHead   string
	uncompletedChar string
	frontText       string
	backText        string
	length          int
}

// StatusCmd represents the album command
var StatusCmd = &cobra.Command{
	Use:   "info",
	Short: "Display general info",
	Long:  ``,
	Run:   status,
}

func status(cmd *cobra.Command, args []string) {
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

	minimumLength := viper.GetInt("appearance.status.bar.minimumlength")
	barLength := findMinInt([]int{len(songNameLine), len(albumNameLine), len(artistNameLine), minimumLength})

	progressStamp := secondsToTimestamp(status.Progress / 1000)
	totalStamp := secondsToTimestamp(fulltrack.Duration / 1000)

	style := styleConfig{
		length:    barLength,
		frontText: progressStamp,
		backText:  totalStamp}

	err = loadStyle(&style)
	if err != nil {
		fmt.Println("Status bar style error:", err)
		return
	}

	fmt.Println(songNameLine)
	fmt.Println(albumNameLine)
	fmt.Println(artistNameLine)

	StaticProgressBar(style, status.Progress, fulltrack.Duration)
}

func loadStyle(style *styleConfig) (err error) {
	startString := fmt.Sprint(viper.Get("appearance.status.bar.startstring"))
	endString := fmt.Sprint(viper.Get("appearance.status.bar.endstring"))
	completedHead := fmt.Sprint(viper.Get("appearance.status.bar.completedhead"))

	completedString := fmt.Sprint(viper.Get("appearance.status.bar.completedchar"))
	uncompletedString := fmt.Sprint(viper.Get("appearance.status.bar.uncompletedchar"))

	if len(completedString) < 1 {
		err := errors.New("completedchar is not of type char")
		return err
	}
	if len(uncompletedString) < 1 {
		err := errors.New("uncompletedchar is not of type char")
		return err
	}

	style.startString = startString
	style.endString = endString
	style.completedChar = completedString
	style.completedHead = completedHead
	style.uncompletedChar = uncompletedString
	return nil
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
	paddingLength := len(style.endString + style.startString + style.frontText + style.backText)
	barLength := style.length - paddingLength
	percentage := int((float32(progress)/float32(total))*float32(barLength) + 0.5)
	fmt.Print(style.frontText + style.startString)

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

	fmt.Print(style.endString + style.backText + "\n")
}

func init() {
}
