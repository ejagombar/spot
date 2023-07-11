package authStore

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
	// "github.com/zmb3/spotify/v2/auth"
	// "golang.org/x/oauth2"
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
	if token.Valid() == true {
		fmt.Println("Token is valid")
	} else {
		fmt.Println("Token is not valid ")
	}
	return nil
}

func GetTime() {
	fmt.Println(time.Now())
}
