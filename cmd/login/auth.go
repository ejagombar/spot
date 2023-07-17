package login

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"strings"
	"time"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"

var (
	//go:embed callback.html
	form  string
	auth  *spotifyauth.Authenticator
	ch    = make(chan *spotify.Client)
	state = "abc123"
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
	client := spotify.New(auth.Client(r.Context(), tok))
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	fmt.Fprintf(w, form)

	ch <- client
}

func getClient() (client *spotify.Client) {
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

	return client
}

func createAuthRequest() error {
	spotify_id := viper.GetString("auth.client_id")
	spotify_client := viper.GetString("auth.client_secret")

	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithClientID(spotify_id),
		spotifyauth.WithClientSecret(spotify_client),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeStreaming,
			spotifyauth.ScopeUserFollowRead,
			spotifyauth.ScopeUserLibraryRead,
			spotifyauth.ScopeUserModifyPlaybackState,
			spotifyauth.ScopePlaylistModifyPrivate,
			spotifyauth.ScopeUserReadPlaybackState,
			spotifyauth.ScopePlaylistModifyPublic,
			spotifyauth.ScopeUserReadCurrentlyPlaying))

	return nil
}

func saveToken(tok *oauth2.Token) error {
	viper.Set("token.access", tok.AccessToken)
	viper.Set("token.refresh", tok.RefreshToken)
	viper.Set("token.timeout", tok.Expiry.Format(time.RFC1123Z))
	viper.WriteConfig()
	return nil
}

func runAuth(cmd *cobra.Command, args []string) error {
	createAuthRequest()
	client := getClient()

	user, err := client.CurrentUser(context.Background())
	if err != nil {
		return err
	}

	fmt.Println(strings.TrimSpace(user.DisplayName) + "'s account connected!")

	tok, err := client.Token()
	if err != nil {
		fmt.Println(err)
	}
	saveToken(tok)

	return nil
}

// authCmd represents the auth command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Connect your spotify account",
	Long:  ``,
	RunE:  runAuth,
}

func init() {
}
