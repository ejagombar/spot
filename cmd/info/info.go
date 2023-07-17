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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	"strconv"
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

// InfoCmd represents the album command
var (
	showAccount bool
	verbose     bool
	barStyle    styleConfig
	InfoCmd     = &cobra.Command{
		Use:   "info",
		Short: "Display general info",
		Long:  ``,
		Run:   info,
	}
)

func info(cmd *cobra.Command, args []string) {
	client, err := authstore.GetClient()
	prechecks.DeviceAvailable(client)
	cobra.CheckErr(err)

	if showAccount {
		accountInfo(client)
	} else {
		songInfo(client)
	}
}

func accountInfo(client *spotify.Client) {
	user, err := client.CurrentUser(context.Background())
	cobra.CheckErr(err)
	devices, err := client.PlayerDevices(context.Background())
	cobra.CheckErr(err)
	player, err := client.PlayerState(context.Background())
	cobra.CheckErr(err)

	bold := color.New(color.Bold).SprintFunc()
	boldRed := color.New(color.FgRed, color.Bold).SprintFunc()

	fmt.Println(boldRed("Account Info"))
	fmt.Println("")
	fmt.Println(bold("Username: ") + user.DisplayName)
	fmt.Println(bold("Followers: ") + strconv.Itoa(int(user.Followers.Count)))
	if player.Playing {
		fmt.Println(bold("Current Device: ") + player.Device.Name)
	}
	if len(devices) > 1 {
		deviceList := devices[0].Name
		for i := 1; i < len(devices); i++ {
			deviceList += ", " + devices[i].Name
		}
		fmt.Println(bold("All available devices: ") + deviceList)
	}
	fmt.Println("")
}

func moreSongInfo(client *spotify.Client) {
	track, err := client.PlayerState(context.Background())
	cobra.CheckErr(err)
	if track.Item == nil {
		return
	}

	songID := track.Item.ID

	fullSong, err := client.GetTrack(context.Background(), songID)
	cobra.CheckErr(err)
	songFeatures, err := client.GetAudioFeatures(context.Background(), songID)
	cobra.CheckErr(err)

	bold := color.New(color.Bold).SprintFunc()
	greenBold := color.New(color.Bold, color.FgGreen).SprintFunc()

	fmt.Println(greenBold("Musical Info"))

	pitch := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

	if i := songFeatures[0].Key % 12; i > 0 {
		fmt.Println(bold("Key: ") + pitch[(i)])
	}
	fmt.Println(bold("Time Signature: ") + fmt.Sprint(songFeatures[0].TimeSignature) + "/4")

	if songFeatures[0].Mode == 0 {
		fmt.Println(bold("Mode: ") + "Minor")
	} else {
		fmt.Println(bold("Mode: ") + "Major")
	}
	fmt.Println(bold("Tempo: ") + fmt.Sprint(int(songFeatures[0].Tempo)) + "bpm")
	fmt.Println("")

	barStyle.length = 40
	barStyle.frontText = ""

	loadStyle(&barStyle)
	fmt.Println(greenBold("Popularity:"))
	StaticProgressBar(&barStyle, fullSong.Popularity, 100)
	fmt.Println(greenBold("Energy:"))
	StaticProgressBar(&barStyle, int(songFeatures[0].Energy*100), 100)
	fmt.Println(greenBold("Valence:"))
	StaticProgressBar(&barStyle, int(songFeatures[0].Valence*100), 100)
	fmt.Println(greenBold("Danceability:"))
	StaticProgressBar(&barStyle, int(songFeatures[0].Danceability*100), 100)
	fmt.Println(greenBold("Loudness:"))
	StaticProgressBar(&barStyle, int(60+songFeatures[0].Loudness), 60) // loudness ranges between -60 and 0
	fmt.Println(greenBold("Instrumentalness:"))
	StaticProgressBar(&barStyle, int(songFeatures[0].Instrumentalness*100), 100)

	fmt.Println("")
}

func songInfo(client *spotify.Client) {
	status, err := client.PlayerCurrentlyPlaying(context.Background())
	cobra.CheckErr(err)
	fulltrack := status.Item
	if fulltrack == nil {
		fmt.Println("No track is playing")
		return
	}
	songTitle := "Song: "
	songField := fulltrack.Name
	albumTitle := "Album: "
	albumField := fulltrack.Album.Name

	var artistTitle string
	artistList := fulltrack.Artists[0].Name

	for i := 1; i < len(fulltrack.Artists); i++ {
		artistList += ", " + fulltrack.Artists[i].Name
	}

	if len(fulltrack.Artists) > 1 {
		artistTitle = "Artists: "
	} else {
		artistTitle = "Artist: "
	}

	minimumLength := viper.GetInt("appearance.status.bar.minimumlength")
	barLength := findMinInt([]int{len(songTitle) + len(songField), len(albumTitle) + len(albumField), len(artistTitle) + len(artistList), minimumLength})

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

	bold := color.New(color.Bold).SprintFunc()
	boldBlue := color.New(color.FgBlue, color.Bold).SprintFunc()

	fmt.Println(boldBlue("Song Info"))

	fmt.Println("")
	fmt.Println(bold(songTitle) + songField)
	fmt.Println(bold(albumTitle) + albumField)
	fmt.Println(bold(artistTitle) + artistList)
	fmt.Println("")

	if verbose {
		moreSongInfo(client)
	}

	color.Set(color.FgBlue)
	StaticProgressBar(&style, status.Progress, fulltrack.Duration)
	color.Unset()
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

func StaticProgressBar(style *styleConfig, progress int, total int) {
	if progress > total {
		progress = total
	}
	if progress < 0 {
		progress = 0
	}
	paddingLength := len([]rune(style.endString + style.startString + style.frontText + style.backText))
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
	InfoCmd.Flags().BoolVar(&showAccount, "account", false, "Displays account instead of song info")
	InfoCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Shows additional song information")
}
