package authStore

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

func GetClient() (client *spotify.Client, err error) {
	client_id := fmt.Sprint(viper.Get("auth.client_id"))
	client_secret := fmt.Sprint(viper.Get("auth.client_secret"))

	if (client_id == "") || (client_secret == "") {
		return nil, errors.New("Client secret or ID missing from config file")
	}

	timeOutStr := fmt.Sprint(viper.Get("token.timeout"))
	access := fmt.Sprint(viper.Get("token.access"))
	refresh := fmt.Sprint(viper.Get("token.refresh"))

	if (timeOutStr == "") || (refresh == "") || (access == "") {
		return nil, errors.New("Token missing from config file")
	}

	auth := spotifyauth.New(
		spotifyauth.WithClientID(client_id),
		spotifyauth.WithClientSecret(client_secret),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeStreaming,
			spotifyauth.ScopeUserFollowRead,
			spotifyauth.ScopeUserModifyPlaybackState))

	timeOut, err := time.Parse(time.RFC1123Z, timeOutStr)
	if err != nil {
		return nil, err
	}

	token := new(oauth2.Token)
	token.AccessToken = access
	token.RefreshToken = refresh
	token.Expiry = timeOut
	ctx := context.Background()
	client = spotify.New(auth.Client(ctx, token))

	return client, nil
}
