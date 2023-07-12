package authStore

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithClientID("19d7b9add81a4e60a5b032a4677c5801"),
		spotifyauth.WithClientSecret("38aa053e6ce44ebc946ec379efd7a7c2"),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeStreaming,
			spotifyauth.ScopeUserFollowRead,
			spotifyauth.ScopeUserModifyPlaybackState))
)

func GetClient() *spotify.Client {
	timeOutStr := fmt.Sprint(viper.Get("token.timeout"))
	access := fmt.Sprint(viper.Get("token.access"))
	refresh := fmt.Sprint(viper.Get("token.refresh"))
	if timeOutStr == "" {
		return nil
	}

	timeOut, err := time.Parse(time.RFC1123Z, timeOutStr)
	if err != nil {
		fmt.Println(timeOut.String())
	}

	token := new(oauth2.Token)
	token.AccessToken = access
	token.RefreshToken = refresh
	token.Expiry = timeOut
	// if token.Valid() == true {
	// fmt.Println("Token is valid ")
	// } else {
	// fmt.Println("Token is not valid ")
	ctx := context.Background()
	client := spotify.New(auth.Client(ctx, token))

	// user, err := client.CurrentUser(ctx)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// }
	// fmt.Println("Logged! in as", user.DisplayName)

	// }
	return client
}

func GetTime() {
	fmt.Println(time.Now())
}
