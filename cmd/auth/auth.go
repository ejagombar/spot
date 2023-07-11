package auth

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"log"
	"net/http"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"

var (
	auth   *spotifyauth.Authenticator
	ch     = make(chan *spotify.Client)
	state  = "abc123"
	client *spotify.Client
	form   = `<!DOCTYPE html>
<html>
<head>
<!-- HTML Codes by Quackit.com -->
<title>
Login Complete!</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
body {background-color:#000000;background-repeat:no-repeat;background-position:top left;background-attachment:fixed;}
h1{text-align:center;font-family:Helvetica, sans-serif;color:#ffffff;background-color:#000000;}
p {text-align:center;font-family:Helvetica, sans-serif;font-size:18px;font-style:normal;font-weight:normal;color:#ffffff;background-color:#000000;}
</style>
</head>
<body>
<h1>Login Complete!</h1>
<p>Please return to your terminal to continue</p>
</body>
</html>
`
)

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client = spotify.New(auth.Client(r.Context(), tok))
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	fmt.Fprintf(w, form)

	ch <- client
}

func getClient() (client *spotify.Client, err error) {
	http.HandleFunc("/callback", completeAuth)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	client = <-ch

	return client, nil
}

func GetClient() *spotify.Client {
	return client
}

func createRequest() error {
	spotify_id := viper.GetString("auth.client_id")
	spotify_client := viper.GetString("auth.client_secret")

	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithClientID(spotify_id),
		spotifyauth.WithClientSecret(spotify_client),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeStreaming,
			spotifyauth.ScopeUserFollowRead,
			spotifyauth.ScopeUserModifyPlaybackState))

	return nil
}

// authCmd represents the auth command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Connect your spotify account",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		createRequest()
		client2, _ := getClient()
		user, _ := client2.CurrentUser(context.Background())
		fmt.Println(user.DisplayName)
		return nil
	},
}

func init() {
}
