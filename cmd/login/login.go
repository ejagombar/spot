package login

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ejagombar/spot/authstore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
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
	state = "1234567IshouldProbablyChangeThis"

	// authCmd represents the auth command
	LoginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login with your spotify account",
		Long: `Login with your spotify account

This step must be done before the CLI program can be used.
Before this command is run, please ensure that the SPOTIFY_CLIENT and SPOTIY_ID fields have been set in the configuration file:

───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Step 1: Go to https://developer.spotify.com/ and login with your spotify account.

Step 2: Go to your dashboard (https://developer.spotify.com/dashboard) and click on the "create an app" button.

Step 3: Enter an app name and description of your choice. (Anything will do)

Step 4: Set the Redirect URL to: http://localhost:8080/callback

Step 5: Click 'create'

Step 6: You will now be greeted with an overview page. At the top right, click "settings"

Step 7: You will see a your client ID and below, a button to reveal your client secret.

Step 8: Copy these values into the spot config file. This hidden file should be found in your home directory with the name .spot.json
    Open the file and enter the client ID and client secret in the appropriate boxes then save and close the file.

Step 9: run the command "spot login". If everything is done correctly, a link will be generated which you can click to login with your spotify account.
    If you need any further help, visit the README or open an issue on the github page: https://github.com/ejagombar/spot 

───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Once the variables are set and this commnd is run, a link will be generated that takes you to a webpage where you can login with your spotify account and allow the tool certian permissions with your account.
The permissions requested are all neccessary in order for the tool to function correctly.`,
		RunE: runAuth,
	}
)

func runAuth(cmd *cobra.Command, args []string) error {
	err := createAuthRequest()
	if err != nil {
		return fmt.Errorf("%w Run 'spot login --help' for instuctions", err)
	}
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
	authstore.SaveToken(tok)

	return nil
}

// Handler function that is used to retrieve the token from the spotify authentication webpage
// This toek is used to create a client.
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

// Starts the callback server, generates a link for the user to login with spotify, and waits
// until a client is recieved which is then returned from the function.
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

// Creates a authentication request with all the nessecary scopes needed for the CLI tool
func createAuthRequest() error {
	spotify_id := viper.GetString("auth.client_id")
	spotify_client := viper.GetString("auth.client_secret")

	if spotify_client == "" || spotify_id == "" {
		return errors.New("CLIENT_ID or CLIENT_SECRET not found.")
	}

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

func init() {
}
