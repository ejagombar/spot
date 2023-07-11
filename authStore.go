package authStore

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	// "github.com/zmb3/spotify/v2/auth"
	// "golang.org/x/oauth2"
)

func getClinet() *spotify.Client {
	timeOut := viper.Get("token.timeout")
	if timeOut != nil {
		return nil
	}

	fmt.Println(time.Now)
	return nil
}
